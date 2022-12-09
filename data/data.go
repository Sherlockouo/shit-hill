package data

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

//WallpaperID is a string representing a wallpaper
type WallpaperID string

//TagID is an int representing a tag on wh
type TagID int64

func (t TagID) String() string {
	return strconv.Itoa(int(t))
}

//Avatar user's avatar
type Avatar struct {
	Two00Px  string `json:"200px"`
	One28Px  string `json:"128px"`
	Three2Px string `json:"32px"`
	Two0Px   string `json:"20px"`
}

//Uploader information on who uploaded a given wallpaper
type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
	Avatar   Avatar `json:"avatar"`
}

//Tag full data on a given wallpaper tag
type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

//Wallpaper information about a given wallpaper
type Wallpaper struct {
	ID         WallpaperID `json:"id"`
	URL        string      `json:"url"`
	ShortURL   string      `json:"short_url"`
	Uploader   Uploader    `json:"uploader"`
	Views      int         `json:"views"`
	Favorites  int         `json:"favorites"`
	Source     string      `json:"source"`
	Purity     string      `json:"purity"`
	Category   string      `json:"category"`
	DimensionX int         `json:"dimension_x"`
	DimensionY int         `json:"dimension_y"`
	Resolution string      `json:"resolution"`
	Ratio      string      `json:"ratio"`
	FileSize   int         `json:"file_size"`
	FileType   string      `json:"file_type"`
	CreatedAt  string      `json:"created_at"`
	Colors     []string    `json:"colors"`
	Path       string      `json:"path"`
	Thumbs     Thumbs      `json:"thumbs"`
	Tags       []Tag       `json:"tags"`
}

// //Download downloads a wallpaper given the local filepath to save the wallpaper to
// func (w *Wallpaper) Download(dir string) error {
// 	resp, err := getAuthedResponse(w.Path)
// 	if err != nil {
// 		return err
// 	}
// 	path := filepath.Join(dir, path.Base(w.Path))
// 	return download(path, resp)
// }

//Thumbs paths for thumbnail images of wallpaper
type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

//TODO: (dlasky) add download support for thumbs

//Meta paging and stats metadata for search results
type Meta struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
	Query       interface{} `json:"query"`
}

//User User preference Data (set on wallhaven's user GUI)
type User struct {
	ThumbSize     string   `json:"thumb_size"`
	PerPage       string   `json:"per_page"`
	Purity        []string `json:"purity"`
	Categories    []string `json:"categories"`
	Resolutions   []string `json:"resolutions"`
	AspectRatios  []string `json:"aspect_ratios"`
	ToplistRange  string   `json:"toplist_range"`
	TagBlacklist  []string `json:"tag_blacklist"`
	UserBlacklist []string `json:"user_blacklist"`
}

//Result Structs -- server responses

//SearchResults a wrapper containing search results from wh
type SearchResults struct {
	Data []Wallpaper `json:"data"`
	Meta Meta        `json:"meta"`
}

//TagResult a wrapper containing a single tag result when TagInfo is requested
type TagResult struct {
	Data Tag `json:"data"`
}

//UserResult a wrapper containing user preference data
type UserResult struct {
	Data User `json:"data"`
}

//WallpaperResult a wrapper containing a single wallpaper result when WallpaperInfo is requested
type WallpaperResult struct {
	Data Wallpaper `json:"data"`
}

// Search Types

//Category is an enum used to represent wallpaper categories
type Category int

//Enum for Category Types
const (
	General Category = 0x100
	Anime   Category = 0x010
	People  Category = 0x001
)

func (c Category) String() string {
	return strconv.FormatInt(int64(c), 2)
}

//Purity is an enum used to represent
type Purity int

//Enum for purity types
const (
	SFW     Purity = 0x100
	Sketchy Purity = 0x010
	NSFW    Purity = 0x001
)

func (p Purity) String() string {
	return strconv.FormatInt(int64(p), 2)
}

//Sort enum specifies the various sort types accepted by WH api
type Sort int

//Sort Enum Values
const (
	DateAdded Sort = iota + 1
	Relevance
	Random
	Views
	Favorites
	Toplist
	Hot
)

func (s Sort) String() string {
	str := [...]string{"", "date_added", "relevance", "random", "views", "favorites", "topList", "hot"}
	return str[s]
}

//Order enum specifies the sort orders accepted by WH api
type Order int

//Sort Enum Values
const (
	Desc Order = iota + 1
	Asc
)

func (o Order) String() string {
	str := [...]string{"", "desc", "asc"}
	return str[o]
}

//TopRange is used to specify the time window for 'top' result when topList is chosen as sort param
type TopRange int

//Enum for TopRange values
const (
	Day TopRange = iota + 1
	ThreeDay
	Week
	Month
	SixMonth
	Year
)

func (t TopRange) String() string {
	str := [...]string{"", "1d", "3d", "1w", "1m", "6m", "1y"}
	return str[t]
}

//Resolution specifies the image resolution to find
type Resolution struct {
	Width  int64
	Height int64
}

func (r Resolution) String() string {
	return fmt.Sprintf("%vx%v", r.Width, r.Height)
}

func (r Resolution) isValid() bool {
	return r.Width > 0 && r.Height > 0
}

//Ratio may be used to specify the aspect ratio of the search
type Ratio struct {
	Horizontal int
	Vertical   int
}

func (r Ratio) String() string {
	return fmt.Sprintf("%vx%v", r.Horizontal, r.Vertical)
}

func (r Ratio) isValid() bool {
	return r.Vertical > 0 && r.Horizontal > 0
}

//Q is used to hold the Q params for various fulltext options that the WH Search supports
type Q struct {
	Tags       []string
	ExcudeTags []string
	UserName   string
	TagID      int
	Type       string //Type is one of png/jpg
	Like       WallpaperID
}

func (q Q) toQuery() url.Values {
	var sb strings.Builder
	for _, tag := range q.Tags {
		sb.WriteString("+")
		sb.WriteString(tag)
	}
	for _, etag := range q.ExcudeTags {
		sb.WriteString("-")
		sb.WriteString(etag)
	}
	if len(q.UserName) > 0 {
		sb.WriteString("@")
		sb.WriteString(q.UserName)
	}
	if len(q.Type) > 0 {
		sb.WriteString("type:")
		sb.WriteString(q.Type)
	}
	out := url.Values{}
	val := sb.String()
	if len(val) > 0 {
		out.Set("q", val)
	}
	return out
}

//Search provides various parameters to search for on wallhaven
type Search struct {
	Query       Q
	Q           string
	Categories  Category
	Purities    Purity
	Sorting     Sort
	Order       Order
	TopRange    TopRange
	AtLeast     Resolution
	Resolutions []Resolution
	Ratios      []Ratio
	Colors      []string //Colors is an array of hex colors represented as strings in #RRGGBB format
	Page        int
}

func (s *Search) ToQuery() url.Values {
	v := s.Query.toQuery()
	if s.Categories > 0 {
		v.Add("categories", s.Categories.String())
	}
	if s.Purities > 0 {
		v.Add("purity", s.Purities.String())
	}
	if s.Sorting > 0 {
		v.Add("sorting", s.Sorting.String())
	}
	if s.Order > 0 {
		v.Add("order", s.Order.String())
	}
	if s.TopRange > 0 && s.Sorting == Toplist {
		v.Add("topRange", s.TopRange.String())
	}
	if s.AtLeast.isValid() {
		v.Add("atleast", s.AtLeast.String())
	}
	if len(s.Resolutions) > 0 {
		outRes := []string{}
		for _, res := range s.Resolutions {
			if res.isValid() {
				outRes = append(outRes, res.String())
			}
		}
		if len(outRes) > 0 {
			v.Add("resolutions", strings.Join(outRes, ","))
		}
	}
	if len(s.Ratios) > 0 {
		outRat := []string{}
		for _, rat := range s.Ratios {
			if rat.isValid() {
				outRat = append(outRat, rat.String())
			}
		}
		if len(outRat) > 0 {
			v.Add("ratios", strings.Join(outRat, ","))
		}
	}
	if len(s.Colors) > 0 {
		v.Add("colors", strings.Join([]string(s.Colors), ","))
	}
	if s.Page > 0 {
		v.Add("page", strconv.Itoa(s.Page))
	}
	return v
}
