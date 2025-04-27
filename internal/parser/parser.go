package parser

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
)

type DJPlaylists struct {
	XMLName    xml.Name   `xml:"DJ_PLAYLISTS"`
	Version    string     `xml:"Version,attr"`
	Product    Product    `xml:"PRODUCT"`
	Collection Collection `xml:"COLLECTION"`
	Playlists  Playlists  `xml:"PLAYLISTS"`
}

type Product struct {
	Name    string `xml:"Name,attr"`
	Version string `xml:"Version,attr"`
	Company string `xml:"Company,attr"`
}

type Collection struct {
	Entries int               `xml:"Entries,attr"`
	Tracks  []CollectionTrack `xml:"TRACK"`
}

// Detailed track with all attributes
type CollectionTrack struct {
	TrackID     string  `xml:"TrackID,attr"`
	Name        string  `xml:"Name,attr"`
	Artist      string  `xml:"Artist,attr"`
	Composer    string  `xml:"Composer,attr"`
	Album       string  `xml:"Album,attr"`
	Grouping    string  `xml:"Grouping,attr"`
	Genre       string  `xml:"Genre,attr"`
	Kind        string  `xml:"Kind,attr"`
	Size        int     `xml:"Size,attr"`
	TotalTime   int     `xml:"TotalTime,attr"`
	DiscNumber  int     `xml:"DiscNumber,attr"`
	TrackNumber int     `xml:"TrackNumber,attr"`
	Year        int     `xml:"Year,attr"`
	AverageBpm  string  `xml:"AverageBpm,attr"`
	DateAdded   string  `xml:"DateAdded,attr"`
	BitRate     int     `xml:"BitRate,attr"`
	SampleRate  int     `xml:"SampleRate,attr"`
	Comments    string  `xml:"Comments,attr"`
	PlayCount   int     `xml:"PlayCount,attr"`
	Rating      int     `xml:"Rating,attr"`
	Location    string  `xml:"Location,attr"`
	Remixer     string  `xml:"Remixer,attr"`
	Tonality    string  `xml:"Tonality,attr"`
	Label       string  `xml:"Label,attr"`
	Mix         string  `xml:"Mix,attr"`
	Tempos      []Tempo `xml:"TEMPO"`
}

// Tempo information nested within tracks
type Tempo struct {
	Inizio  string `xml:"Inizio,attr"`
	Bpm     string `xml:"Bpm,attr"`
	Metro   string `xml:"Metro,attr"`
	Battito string `xml:"Battito,attr"`
}

// Playlists section
type Playlists struct {
	RootNode Node `xml:"NODE"`
}

// Node - can be either a folder or playlist
type Node struct {
	Type     string          `xml:"Type,attr"`
	Name     string          `xml:"Name,attr"`
	Count    int             `xml:"Count,attr,omitempty"`
	KeyType  string          `xml:"KeyType,attr,omitempty"`
	Entries  int             `xml:"Entries,attr,omitempty"`
	Tracks   []PlaylistTrack `xml:"TRACK,omitempty"`
	Children []Node          `xml:"NODE,omitempty"`
}

// The simplified track reference used in playlists
type PlaylistTrack struct {
	Key string `xml:"Key,attr"`
}

func (p DJPlaylists) String() string {
	return fmt.Sprintf("DJPlaylists{Version: %s, Product: %s, Collection: %d tracks, Playlists: %d playlists}",
		p.Version, p.Product.Name, len(p.Collection.Tracks), len(p.Playlists.RootNode.Children))
}

func (t CollectionTrack) String() string {
	return fmt.Sprintf("CollectionTrack{TrackID: %s, Name: %s, Artist: %s, Album: %s, Genre: %s, BPM: %s}",
		t.TrackID, t.Name, t.Artist, t.Album, t.Genre, t.AverageBpm)
}

func (p Playlists) String() string {
	var childrenStrings string
	for _, child := range p.RootNode.Children {
		childrenStrings += fmt.Sprintln(child)
	}
	return fmt.Sprintf("PLAYLISTS{RootNode: %s, Count: %d, Entries: %d, Children: {\n%s}\n}", p.RootNode.Name, p.RootNode.Count, p.RootNode.Entries, childrenStrings)
}

func (n Node) String() string {
	if n.Type == "1" {
		var playlistStrings string
		for _, child := range n.Tracks {
			playlistStrings += fmt.Sprintln("{Playlist: key:", child.Key, " }")
		}
		return fmt.Sprintf("PLAYLIST{Name: %s, Count: %d, Tracks: {\n%s}", n.Name, n.Count, playlistStrings)
	} else {

		var folderStrings string
		for _, child := range n.Children {
			folderStrings += fmt.Sprintln(child)

		}
		return fmt.Sprintf("FOLDER{Name: %s, Count: %d, Entries: %d, Children: {\n%s}", n.Name, n.Count, n.Entries, folderStrings)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExtractFilePath(fileURL string) (string, error) {
	// Parse the URL
	u, err := url.Parse(fileURL)
	check(err)

	// Get the path from the URL and convert it to local path format
	path := u.Path

	// On Windows, remove the leading slash and fix drive letter format
	if strings.HasPrefix(path, "/") {
		path = path[1:] // Remove leading slash
	}

	// URL decode the path (convert %20 to spaces, etc.)
	path, err = url.PathUnescape(path)
	check(err)

	return path, nil
}

func PrintAllParsedFile(playlists DJPlaylists) {
	fmt.Println(playlists)
	fmt.Printf("{ \nCollection:, Count: %d, Tracks: {\n", len(playlists.Collection.Tracks))
	for _, track := range playlists.Collection.Tracks {
		fmt.Println(track)
	}
	fmt.Println("} \nPlaylists: {")
	fmt.Println(playlists.Playlists)
}
