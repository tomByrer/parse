// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT
package css

// github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go
type Hash uint32

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// Hash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := _Hash_fnv(s)
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	return 0
}

// _Hash_fnv computes the FNV hash with an arbitrary starting value h.
func _Hash_fnv(s []byte) uint32 {
	h := uint32(_Hash_hash0)
	for i := range s {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func _Hash_match(s string, t []byte) bool {
	for i, c := range t {
		if s[i] != c {
			return false
		}
	}
	return true
}

func _Hash_string(i Hash) string {
	return _Hash_text[i>>8 : i>>8+i&0xff]
}

const _Hash_hash0 = 0x700e0976
const _Hash_maxLen = 27
const _Hash_text = "background-position-ybackground-repeatext-justifybehaviorgba" +
	"ckground-attachmentext-kashida-spacempty-cellsans-serifantas" +
	"yblackblanchedalmondarkblueboldarkcyanborder-bottom-colorich" +
	"nesscrollbar-3d-light-coloroyalblueborder-bottom-stylemonchi" +
	"ffont-familyborder-bottom-widthslavenderblushborder-collapse" +
	"ashellayer-background-coloruby-alignborder-coloruby-overhang" +
	"ainsborosybrownborder-left-coloruby-positionborder-left-styl" +
	"eborder-left-widthborder-right-colorborder-right-styleborder" +
	"-right-widthborder-spacinghostwhitext-autospaceborder-styleb" +
	"order-top-colorborder-top-styleborder-top-widthborder-widthb" +
	"urlywoodarkgoldenrodarkgraycadetbluecaption-sideeppinkchocol" +
	"atext-align-lastresscrollbar-arrow-colorclearclipadding-bott" +
	"omargin-bottomargin-leftext-overflow-xcontentext-shadoword-b" +
	"reakcornflowerbluecornsilkcounter-incrementext-transformargi" +
	"n-rightext-underline-positioncounter-resetcue-aftercue-befor" +
	"estgreencursivertical-aligncursordarkslatebluedarkslategrayd" +
	"arkturquoisedarkvioletdisplay-duringdodgerbluefirebrickfloat" +
	"floralwhitesmokefont-size-adjustfont-stretchartreusefont-sty" +
	"lefont-variantiquewhite-spacefont-weightfuchsiacceleratorpha" +
	"nsaddlebrownavajowhitext-decorationonelevationormalawngreeni" +
	"mportantindianredarkmagentable-layout-floword-spacinginherit" +
	"initialicebluevioletter-spacinglayout-grid-char-spacinglayou" +
	"t-grid-line-breaklayout-grid-modefaultlayout-grid-typeachpuf" +
	"filterlightbluelightcoralphazimuthoneydewidowsandybrownlight" +
	"cyanlightgoldenrodyelloword-wrapadding-leftlightgraylightgre" +
	"enlightpinklightsalmonlightseagreenlightskybluelightslateblu" +
	"elightsteelbluelightyellowriting-modelimegreenline-heightlis" +
	"t-style-imagelist-style-positionlist-style-typemarker-offset" +
	"markscrollbar-base-colormax-heightmax-widthmediumbluemediumo" +
	"rchidarkolivegreenmediumpurplemediumseagreenmediumslatebluem" +
	"ediumspringgreenmediumturquoisemediumvioletredarkorangeredar" +
	"kgreenyellowgreenmidnightbluemin-heightmin-widthmintcreamarg" +
	"in-topadding-rightmistyrosemoccasinclude-sourceolivedrabackg" +
	"round-position-xoutline-coloroutline-styleoutline-widthoverf" +
	"low-ypage-break-beforepage-break-insidepalegoldenrodarkorchi" +
	"darkkhakime-modeepskybluepalegreenpaleturquoisepalevioletred" +
	"arksalmonospacepapayawhipadding-topage-break-afterpause-afte" +
	"rpause-beforepitch-rangepowderblueprogidarkseagreenquotescro" +
	"llbar-dark-shadow-colorscrollbar-face-colorscrollbar-highlig" +
	"ht-colorscrollbar-shadow-colorscrollbar-track-colorspeak-hea" +
	"derspeak-numeralayer-background-imagespeak-punctuationspeech" +
	"-ratext-indentunicode-bidirectionvisibilityvoice-familyvolum" +
	"ediumaquamarine"

const (
	Accelerator                 Hash = 0x4660b
	Aliceblue                   Hash = 0x4f109
	Alpha                       Hash = 0x57b05
	Antiquewhite                Hash = 0x4430c
	Aquamarine                  Hash = 0x9dd0a
	Azimuth                     Hash = 0x57f07
	Background                  Hash = 0xa
	Background_Attachment       Hash = 0x3a15
	Background_Color            Hash = 0x13710
	Background_Image            Hash = 0x97510
	Background_Position         Hash = 0x13
	Background_Position_X       Hash = 0x7b715
	Background_Position_Y       Hash = 0x15
	Background_Repeat           Hash = 0x1511
	Behavior                    Hash = 0x3108
	Black                       Hash = 0x7905
	Blanchedalmond              Hash = 0x7e0e
	Blueviolet                  Hash = 0x4f60a
	Bold                        Hash = 0x9304
	Border                      Hash = 0x9e06
	Border_Bottom               Hash = 0x9e0d
	Border_Bottom_Color         Hash = 0x9e13
	Border_Bottom_Style         Hash = 0xd713
	Border_Bottom_Width         Hash = 0xfc13
	Border_Collapse             Hash = 0x11d0f
	Border_Color                Hash = 0x1500c
	Border_Left                 Hash = 0x1770b
	Border_Left_Color           Hash = 0x17711
	Border_Left_Style           Hash = 0x19411
	Border_Left_Width           Hash = 0x1a511
	Border_Right                Hash = 0x1b60c
	Border_Right_Color          Hash = 0x1b612
	Border_Right_Style          Hash = 0x1c812
	Border_Right_Width          Hash = 0x1da12
	Border_Spacing              Hash = 0x1ec0e
	Border_Style                Hash = 0x20f0c
	Border_Top                  Hash = 0x21b0a
	Border_Top_Color            Hash = 0x21b10
	Border_Top_Style            Hash = 0x22b10
	Border_Top_Width            Hash = 0x23b10
	Border_Width                Hash = 0x24b0c
	Bottom                      Hash = 0xa506
	Burlywood                   Hash = 0x25709
	Cadetblue                   Hash = 0x27309
	Caption_Side                Hash = 0x27c0c
	Chartreuse                  Hash = 0x4260a
	Chocolate                   Hash = 0x28e09
	Clear                       Hash = 0x2bc05
	Clip                        Hash = 0x2c104
	Color                       Hash = 0xac05
	Content                     Hash = 0x2f607
	Cornflowerblue              Hash = 0x3100e
	Cornsilk                    Hash = 0x31e08
	Counter_Increment           Hash = 0x32611
	Counter_Reset               Hash = 0x3650d
	Cue                         Hash = 0x37203
	Cue_After                   Hash = 0x37209
	Cue_Before                  Hash = 0x37b0a
	Cursive                     Hash = 0x38c07
	Cursor                      Hash = 0x39f06
	Darkblue                    Hash = 0x8b08
	Darkcyan                    Hash = 0x9608
	Darkgoldenrod               Hash = 0x25f0d
	Darkgray                    Hash = 0x26b08
	Darkgreen                   Hash = 0x74109
	Darkkhaki                   Hash = 0x83409
	Darkmagenta                 Hash = 0x4c00b
	Darkolivegreen              Hash = 0x6d00e
	Darkorange                  Hash = 0x7350a
	Darkorchid                  Hash = 0x82b0a
	Darksalmon                  Hash = 0x86f0a
	Darkseagreen                Hash = 0x8d30c
	Darkslateblue               Hash = 0x3a50d
	Darkslategray               Hash = 0x3b20d
	Darkturquoise               Hash = 0x3bf0d
	Darkviolet                  Hash = 0x3cc0a
	Deeppink                    Hash = 0x28608
	Deepskyblue                 Hash = 0x8420b
	Default                     Hash = 0x54707
	Direction                   Hash = 0x9b409
	Display                     Hash = 0x3d607
	Dodgerblue                  Hash = 0x3e40a
	Elevation                   Hash = 0x49909
	Empty_Cells                 Hash = 0x5f0b
	Fantasy                     Hash = 0x7207
	Filter                      Hash = 0x56406
	Firebrick                   Hash = 0x3ee09
	Float                       Hash = 0x3f705
	Floralwhite                 Hash = 0x3fc0b
	Font                        Hash = 0xf104
	Font_Family                 Hash = 0xf10b
	Font_Size                   Hash = 0x40c09
	Font_Size_Adjust            Hash = 0x40c10
	Font_Stretch                Hash = 0x41c0c
	Font_Style                  Hash = 0x4300a
	Font_Variant                Hash = 0x43a0c
	Font_Weight                 Hash = 0x4550b
	Forestgreen                 Hash = 0x3810b
	Fuchsia                     Hash = 0x46007
	Gainsboro                   Hash = 0x16709
	Ghostwhite                  Hash = 0x1f90a
	Goldenrod                   Hash = 0x26309
	Greenyellow                 Hash = 0x7450b
	Height                      Hash = 0x64b06
	Honeydew                    Hash = 0x58508
	Hsl                         Hash = 0x10e03
	Hsla                        Hash = 0x10e04
	Ime_Mode                    Hash = 0x83c08
	Important                   Hash = 0x4af09
	Include_Source              Hash = 0x7a10e
	Indianred                   Hash = 0x4b809
	Inherit                     Hash = 0x4e507
	Initial                     Hash = 0x4ec07
	Lavender                    Hash = 0x11008
	Lavenderblush               Hash = 0x1100d
	Lawngreen                   Hash = 0x4a609
	Layer_Background_Color      Hash = 0x13116
	Layer_Background_Image      Hash = 0x96f16
	Layout_Flow                 Hash = 0x4cf0b
	Layout_Grid                 Hash = 0x50b0b
	Layout_Grid_Char            Hash = 0x50b10
	Layout_Grid_Char_Spacing    Hash = 0x50b18
	Layout_Grid_Line            Hash = 0x52310
	Layout_Grid_Mode            Hash = 0x53910
	Layout_Grid_Type            Hash = 0x54e10
	Left                        Hash = 0x17e04
	Lemonchiffon                Hash = 0xe80c
	Letter_Spacing              Hash = 0x4fd0e
	Lightblue                   Hash = 0x56a09
	Lightcoral                  Hash = 0x5730a
	Lightcyan                   Hash = 0x59b09
	Lightgoldenrodyellow        Hash = 0x5a414
	Lightgray                   Hash = 0x5cb09
	Lightgreen                  Hash = 0x5d40a
	Lightpink                   Hash = 0x5de09
	Lightsalmon                 Hash = 0x5e70b
	Lightseagreen               Hash = 0x5f20d
	Lightskyblue                Hash = 0x5ff0c
	Lightslateblue              Hash = 0x60b0e
	Lightsteelblue              Hash = 0x6190e
	Lightyellow                 Hash = 0x6270b
	Limegreen                   Hash = 0x63d09
	Line_Break                  Hash = 0x52f0a
	Line_Height                 Hash = 0x6460b
	List_Style                  Hash = 0x6510a
	List_Style_Image            Hash = 0x65110
	List_Style_Position         Hash = 0x66113
	List_Style_Type             Hash = 0x6740f
	Magenta                     Hash = 0x4c407
	Margin                      Hash = 0x2d106
	Margin_Bottom               Hash = 0x2d10d
	Margin_Left                 Hash = 0x2dd0b
	Margin_Right                Hash = 0x3430c
	Margin_Top                  Hash = 0x77c0a
	Marker_Offset               Hash = 0x6830d
	Marks                       Hash = 0x69005
	Max_Height                  Hash = 0x6a80a
	Max_Width                   Hash = 0x6b209
	Mediumaquamarine            Hash = 0x9d710
	Mediumblue                  Hash = 0x6bb0a
	Mediumorchid                Hash = 0x6c50c
	Mediumpurple                Hash = 0x6de0c
	Mediumseagreen              Hash = 0x6ea0e
	Mediumslateblue             Hash = 0x6f80f
	Mediumspringgreen           Hash = 0x70711
	Mediumturquoise             Hash = 0x7180f
	Mediumvioletred             Hash = 0x7270f
	Midnightblue                Hash = 0x7550c
	Min_Height                  Hash = 0x7610a
	Min_Width                   Hash = 0x76b09
	Mintcream                   Hash = 0x77409
	Mistyrose                   Hash = 0x79209
	Moccasin                    Hash = 0x79b08
	Monospace                   Hash = 0x87609
	Navajowhite                 Hash = 0x47f0b
	None                        Hash = 0x49604
	Normal                      Hash = 0x4a106
	Olivedrab                   Hash = 0x7af09
	Orangered                   Hash = 0x73909
	Orphans                     Hash = 0x46f07
	Outline                     Hash = 0x7cc07
	Outline_Color               Hash = 0x7cc0d
	Outline_Style               Hash = 0x7d90d
	Outline_Width               Hash = 0x7e60d
	Overflow                    Hash = 0x2ec08
	Overflow_X                  Hash = 0x2ec0a
	Overflow_Y                  Hash = 0x7f30a
	Padding                     Hash = 0x2c407
	Padding_Bottom              Hash = 0x2c40e
	Padding_Left                Hash = 0x5bf0c
	Padding_Right               Hash = 0x7850d
	Padding_Top                 Hash = 0x8880b
	Page                        Hash = 0x7fd04
	Page_Break_After            Hash = 0x89210
	Page_Break_Before           Hash = 0x7fd11
	Page_Break_Inside           Hash = 0x80e11
	Palegoldenrod               Hash = 0x81f0d
	Palegreen                   Hash = 0x84d09
	Paleturquoise               Hash = 0x8560d
	Palevioletred               Hash = 0x8630d
	Papayawhip                  Hash = 0x87f0a
	Pause                       Hash = 0x8a205
	Pause_After                 Hash = 0x8a20b
	Pause_Before                Hash = 0x8ad0c
	Peachpuff                   Hash = 0x55c09
	Pitch                       Hash = 0x8b905
	Pitch_Range                 Hash = 0x8b90b
	Play_During                 Hash = 0x3d90b
	Position                    Hash = 0xb08
	Powderblue                  Hash = 0x8c40a
	Progid                      Hash = 0x8ce06
	Quotes                      Hash = 0x8df06
	Rgb                         Hash = 0x3803
	Rgba                        Hash = 0x3804
	Richness                    Hash = 0xb008
	Right                       Hash = 0x1bd05
	Rosybrown                   Hash = 0x16e09
	Royalblue                   Hash = 0xce09
	Ruby_Align                  Hash = 0x1460a
	Ruby_Overhang               Hash = 0x15b0d
	Ruby_Position               Hash = 0x1870d
	Saddlebrown                 Hash = 0x4750b
	Sandybrown                  Hash = 0x5910a
	Sans_Serif                  Hash = 0x690a
	Scrollbar_3d_Light_Color    Hash = 0xb718
	Scrollbar_Arrow_Color       Hash = 0x2a715
	Scrollbar_Base_Color        Hash = 0x69414
	Scrollbar_Dark_Shadow_Color Hash = 0x8e41b
	Scrollbar_Face_Color        Hash = 0x8ff14
	Scrollbar_Highlight_Color   Hash = 0x91319
	Scrollbar_Shadow_Color      Hash = 0x92c16
	Scrollbar_Track_Color       Hash = 0x94215
	Seagreen                    Hash = 0x5f708
	Seashell                    Hash = 0x12a08
	Serif                       Hash = 0x6e05
	Size                        Hash = 0x41104
	Slateblue                   Hash = 0x3a909
	Slategray                   Hash = 0x3b609
	Speak                       Hash = 0x95705
	Speak_Header                Hash = 0x9570c
	Speak_Numeral               Hash = 0x9630d
	Speak_Punctuation           Hash = 0x98511
	Speech_Rate                 Hash = 0x9960b
	Springgreen                 Hash = 0x70d0b
	Steelblue                   Hash = 0x61e09
	Stress                      Hash = 0x2a206
	Table_Layout                Hash = 0x4c90c
	Text_Align                  Hash = 0x2950a
	Text_Align_Last             Hash = 0x2950f
	Text_Autospace              Hash = 0x2010e
	Text_Decoration             Hash = 0x4880f
	Text_Indent                 Hash = 0x99f0b
	Text_Justify                Hash = 0x250c
	Text_Kashida_Space          Hash = 0x4e12
	Text_Overflow               Hash = 0x2e70d
	Text_Shadow                 Hash = 0x2fc0b
	Text_Transform              Hash = 0x3360e
	Text_Underline_Position     Hash = 0x34e17
	Top                         Hash = 0x22203
	Turquoise                   Hash = 0x3c309
	Unicode_Bidi                Hash = 0x9aa0c
	Vertical_Align              Hash = 0x3910e
	Visibility                  Hash = 0x9bd0a
	Voice_Family                Hash = 0x9c70c
	Volume                      Hash = 0x9d306
	White                       Hash = 0x1fe05
	White_Space                 Hash = 0x44a0b
	Whitesmoke                  Hash = 0x4020a
	Widows                      Hash = 0x58c06
	Width                       Hash = 0x10a05
	Word_Break                  Hash = 0x3060a
	Word_Spacing                Hash = 0x4d90c
	Word_Wrap                   Hash = 0x5b709
	Writing_Mode                Hash = 0x6310c
	Yellow                      Hash = 0x5b206
	Yellowgreen                 Hash = 0x74a0b
)

var _Hash_table = [1 << 9]Hash{
	0x0:   0x5910a, // sandybrown
	0x1:   0x22203, // top
	0x4:   0xce09,  // royalblue
	0x6:   0x4880f, // text-decoration
	0xb:   0x4cf0b, // layout-flow
	0xc:   0x13710, // background-color
	0xd:   0xa506,  // bottom
	0x10:  0x5f20d, // lightseagreen
	0x11:  0x8420b, // deepskyblue
	0x12:  0x3a909, // slateblue
	0x13:  0x5f0b,  // empty-cells
	0x14:  0x2c104, // clip
	0x15:  0x6bb0a, // mediumblue
	0x18:  0x2d10d, // margin-bottom
	0x1a:  0x1500c, // border-color
	0x1b:  0x58508, // honeydew
	0x1d:  0x24b0c, // border-width
	0x1e:  0x9570c, // speak-header
	0x1f:  0x8630d, // palevioletred
	0x20:  0x1ec0e, // border-spacing
	0x22:  0x2c407, // padding
	0x23:  0x3430c, // margin-right
	0x27:  0x76b09, // min-width
	0x29:  0x5cb09, // lightgray
	0x2a:  0x6270b, // lightyellow
	0x2c:  0x89210, // page-break-after
	0x2d:  0x2f607, // content
	0x30:  0x250c,  // text-justify
	0x32:  0x2950f, // text-align-last
	0x34:  0x8ff14, // scrollbar-face-color
	0x37:  0x4c407, // magenta
	0x38:  0x3b609, // slategray
	0x3a:  0x97510, // background-image
	0x3c:  0x7a10e, // include-source
	0x3d:  0x61e09, // steelblue
	0x3e:  0x4fd0e, // letter-spacing
	0x40:  0x11d0f, // border-collapse
	0x41:  0x11008, // lavender
	0x44:  0x6460b, // line-height
	0x45:  0x98511, // speak-punctuation
	0x46:  0x9bd0a, // visibility
	0x47:  0x2bc05, // clear
	0x4b:  0x4f60a, // blueviolet
	0x4e:  0x54707, // default
	0x50:  0x6830d, // marker-offset
	0x52:  0x32611, // counter-increment
	0x53:  0x60b0e, // lightslateblue
	0x54:  0x12a08, // seashell
	0x56:  0x1870d, // ruby-position
	0x57:  0x7d90d, // outline-style
	0x58:  0x5f708, // seagreen
	0x59:  0xac05,  // color
	0x5c:  0x27c0c, // caption-side
	0x5d:  0x64b06, // height
	0x5e:  0x6f80f, // mediumslateblue
	0x5f:  0x8ad0c, // pause-before
	0x60:  0xe80c,  // lemonchiffon
	0x63:  0x38c07, // cursive
	0x66:  0x47f0b, // navajowhite
	0x67:  0x9c70c, // voice-family
	0x68:  0x25f0d, // darkgoldenrod
	0x69:  0x3e40a, // dodgerblue
	0x6a:  0x4300a, // font-style
	0x6b:  0x9b409, // direction
	0x6d:  0x7350a, // darkorange
	0x6f:  0x43a0c, // font-variant
	0x70:  0x2d106, // margin
	0x71:  0x7fd11, // page-break-before
	0x73:  0x2e70d, // text-overflow
	0x74:  0x4e12,  // text-kashida-space
	0x75:  0x31e08, // cornsilk
	0x76:  0x4550b, // font-weight
	0x77:  0x41104, // size
	0x78:  0x50b0b, // layout-grid
	0x79:  0x8880b, // padding-top
	0x7d:  0x79209, // mistyrose
	0x7e:  0x57f07, // azimuth
	0x7f:  0x8a20b, // pause-after
	0x84:  0x39f06, // cursor
	0x85:  0x10e03, // hsl
	0x86:  0x77409, // mintcream
	0x8b:  0x3ee09, // firebrick
	0x8d:  0x37209, // cue-after
	0x8f:  0x37b0a, // cue-before
	0x91:  0x7207,  // fantasy
	0x94:  0x15b0d, // ruby-overhang
	0x95:  0x2c40e, // padding-bottom
	0x9a:  0x56a09, // lightblue
	0x9c:  0x86f0a, // darksalmon
	0x9d:  0x40c10, // font-size-adjust
	0x9e:  0x1f90a, // ghostwhite
	0xa0:  0x8d30c, // darkseagreen
	0xa2:  0x80e11, // page-break-inside
	0xa4:  0x26309, // goldenrod
	0xa6:  0x4d90c, // word-spacing
	0xa7:  0x50b18, // layout-grid-char-spacing
	0xa9:  0x4af09, // important
	0xaa:  0x7610a, // min-height
	0xb0:  0x17711, // border-left-color
	0xb1:  0x7fd04, // page
	0xb2:  0x96f16, // layer-background-image
	0xb5:  0x52310, // layout-grid-line
	0xb6:  0x1511,  // background-repeat
	0xb7:  0x9e13,  // border-bottom-color
	0xb9:  0x26b08, // darkgray
	0xbb:  0x5bf0c, // padding-left
	0xbc:  0x1bd05, // right
	0xc0:  0x69414, // scrollbar-base-color
	0xc1:  0x6190e, // lightsteelblue
	0xc2:  0x10a05, // width
	0xc5:  0x3c309, // turquoise
	0xc8:  0x3f705, // float
	0xca:  0x1460a, // ruby-align
	0xcb:  0xb08,   // position
	0xcc:  0x77c0a, // margin-top
	0xce:  0x2dd0b, // margin-left
	0xcf:  0x2fc0b, // text-shadow
	0xd0:  0x3060a, // word-break
	0xd4:  0x4020a, // whitesmoke
	0xd6:  0x34e17, // text-underline-position
	0xd7:  0x1da12, // border-right-width
	0xd8:  0x7af09, // olivedrab
	0xd9:  0x84d09, // palegreen
	0xdc:  0x69005, // marks
	0xdd:  0x3cc0a, // darkviolet
	0xde:  0x13,    // background-position
	0xe0:  0x9d710, // mediumaquamarine
	0xe1:  0x9304,  // bold
	0xe2:  0x7180f, // mediumturquoise
	0xe4:  0x81f0d, // palegoldenrod
	0xe5:  0x4c00b, // darkmagenta
	0xe6:  0x16e09, // rosybrown
	0xe7:  0x1a511, // border-left-width
	0xe8:  0x83409, // darkkhaki
	0xea:  0x7e0e,  // blanchedalmond
	0xeb:  0x4ec07, // initial
	0xec:  0x10e04, // hsla
	0xee:  0x4750b, // saddlebrown
	0xef:  0x8560d, // paleturquoise
	0xf1:  0x1b612, // border-right-color
	0xf3:  0x1fe05, // white
	0xf7:  0x91319, // scrollbar-highlight-color
	0xf9:  0x53910, // layout-grid-mode
	0xfc:  0x20f0c, // border-style
	0xfe:  0x66113, // list-style-position
	0x100: 0x13116, // layer-background-color
	0x102: 0x54e10, // layout-grid-type
	0x103: 0x1770b, // border-left
	0x104: 0x2ec08, // overflow
	0x105: 0x7550c, // midnightblue
	0x10b: 0x2950a, // text-align
	0x10e: 0x22b10, // border-top-style
	0x110: 0x5a414, // lightgoldenrodyellow
	0x114: 0x9e06,  // border
	0x119: 0xf104,  // font
	0x11c: 0x9dd0a, // aquamarine
	0x11d: 0x5d40a, // lightgreen
	0x11e: 0x5b206, // yellow
	0x120: 0x95705, // speak
	0x121: 0x44a0b, // white-space
	0x123: 0x3a50d, // darkslateblue
	0x125: 0x2010e, // text-autospace
	0x128: 0x1100d, // lavenderblush
	0x12c: 0x5e70b, // lightsalmon
	0x12d: 0x4e507, // inherit
	0x131: 0x82b0a, // darkorchid
	0x132: 0x21b0a, // border-top
	0x133: 0x3d90b, // play-during
	0x137: 0x23b10, // border-top-width
	0x139: 0x46f07, // orphans
	0x13a: 0xf10b,  // font-family
	0x13f: 0x87f0a, // papayawhip
	0x140: 0x8a205, // pause
	0x143: 0x3100e, // cornflowerblue
	0x144: 0x3d607, // display
	0x146: 0x4f109, // aliceblue
	0x14a: 0x8b08,  // darkblue
	0x14b: 0x3108,  // behavior
	0x14c: 0x3650d, // counter-reset
	0x14d: 0x7450b, // greenyellow
	0x14e: 0x70711, // mediumspringgreen
	0x14f: 0x8c40a, // powderblue
	0x150: 0x50b10, // layout-grid-char
	0x158: 0x7cc07, // outline
	0x159: 0x25709, // burlywood
	0x15b: 0xfc13,  // border-bottom-width
	0x15c: 0x49604, // none
	0x15e: 0x37203, // cue
	0x15f: 0x4c90c, // table-layout
	0x160: 0x8b90b, // pitch-range
	0x162: 0x2a206, // stress
	0x163: 0x7b715, // background-position-x
	0x165: 0x4a106, // normal
	0x167: 0x6de0c, // mediumpurple
	0x169: 0x5730a, // lightcoral
	0x16c: 0x6a80a, // max-height
	0x16d: 0x3804,  // rgba
	0x16e: 0x65110, // list-style-image
	0x170: 0x28608, // deeppink
	0x173: 0x8ce06, // progid
	0x175: 0x70d0b, // springgreen
	0x176: 0x3810b, // forestgreen
	0x179: 0x79b08, // moccasin
	0x17a: 0x7270f, // mediumvioletred
	0x17e: 0x99f0b, // text-indent
	0x181: 0x6740f, // list-style-type
	0x182: 0x16709, // gainsboro
	0x183: 0x3bf0d, // darkturquoise
	0x184: 0x3b20d, // darkslategray
	0x189: 0x2ec0a, // overflow-x
	0x18b: 0x8df06, // quotes
	0x18c: 0x3a15,  // background-attachment
	0x18f: 0x1b60c, // border-right
	0x191: 0x7905,  // black
	0x192: 0x74a0b, // yellowgreen
	0x194: 0x55c09, // peachpuff
	0x197: 0x3fc0b, // floralwhite
	0x19c: 0x6d00e, // darkolivegreen
	0x19d: 0x5b709, // word-wrap
	0x19e: 0x19411, // border-left-style
	0x1a0: 0x9960b, // speech-rate
	0x1a1: 0x7e60d, // outline-width
	0x1a2: 0x9aa0c, // unicode-bidi
	0x1a3: 0x6510a, // list-style
	0x1a4: 0x8b905, // pitch
	0x1a5: 0x94215, // scrollbar-track-color
	0x1a6: 0x46007, // fuchsia
	0x1a8: 0x3910e, // vertical-align
	0x1ad: 0x57b05, // alpha
	0x1ae: 0x6b209, // max-width
	0x1af: 0xb008,  // richness
	0x1b0: 0x3803,  // rgb
	0x1b1: 0x7850d, // padding-right
	0x1b2: 0x2a715, // scrollbar-arrow-color
	0x1b3: 0x17e04, // left
	0x1b5: 0x49909, // elevation
	0x1b6: 0x52f0a, // line-break
	0x1ba: 0x28e09, // chocolate
	0x1bb: 0x9630d, // speak-numeral
	0x1bd: 0x4660b, // accelerator
	0x1be: 0x63d09, // limegreen
	0x1c1: 0x9608,  // darkcyan
	0x1c3: 0x5ff0c, // lightskyblue
	0x1c5: 0x690a,  // sans-serif
	0x1c6: 0x9e0d,  // border-bottom
	0x1c7: 0xa,     // background
	0x1c8: 0x9d306, // volume
	0x1ca: 0x6310c, // writing-mode
	0x1cb: 0xb718,  // scrollbar-3d-light-color
	0x1cc: 0x58c06, // widows
	0x1cf: 0x40c09, // font-size
	0x1d0: 0x15,    // background-position-y
	0x1d1: 0x59b09, // lightcyan
	0x1d4: 0x4b809, // indianred
	0x1db: 0x73909, // orangered
	0x1dc: 0x4430c, // antiquewhite
	0x1dd: 0x4a609, // lawngreen
	0x1df: 0x6ea0e, // mediumseagreen
	0x1e0: 0x21b10, // border-top-color
	0x1e2: 0x5de09, // lightpink
	0x1e4: 0x3360e, // text-transform
	0x1e6: 0x6c50c, // mediumorchid
	0x1e9: 0x87609, // monospace
	0x1ec: 0x92c16, // scrollbar-shadow-color
	0x1ed: 0x74109, // darkgreen
	0x1ef: 0x27309, // cadetblue
	0x1f0: 0x56406, // filter
	0x1f1: 0x1c812, // border-right-style
	0x1f6: 0x7f30a, // overflow-y
	0x1f7: 0x7cc0d, // outline-color
	0x1fa: 0xd713,  // border-bottom-style
	0x1fb: 0x41c0c, // font-stretch
	0x1fc: 0x8e41b, // scrollbar-dark-shadow-color
	0x1fd: 0x83c08, // ime-mode
	0x1fe: 0x4260a, // chartreuse
	0x1ff: 0x6e05,  // serif
}