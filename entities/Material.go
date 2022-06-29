package entities

type File struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Description string  `json:"description"`
    FileType string `json:"file_type"`
    Category string `json:"category"`
    Tags  []string `json:"tags"`
    Classification string `json:"classification"`
    Link string `json:"link"`
    Location string `json:"location"`
    Author string `json:"author"`
    BibleReferences []BibleReference `json:"bible_reference"`
    Transcripts string `json:"transcripts"`
    Language string `json:"language"`
    CreationDate string `json:"creation_date"`
}

type BibleReference struct {
    BibleBook string `json:"bible_book"`
    BibleBookNumber int32 `json:"bible_book_number"` // 1 Corinthians, 2 Peter, 3 John
    BibleChapter int32 `json:"bible_chapter"`
    BibleVerse []int32 `json:"bible_verse"`
    BibleVersion string `json:"bible_version"`
}