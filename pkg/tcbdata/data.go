package tcbdata

import (
	"github.com/krau/SaveAny-Bot/pkg/enums/tasktype"
	"github.com/krau/SaveAny-Bot/pkg/parser"
	"github.com/krau/SaveAny-Bot/pkg/telegraph"
	"github.com/krau/SaveAny-Bot/pkg/tfile"
)

const (
	TypeAdd        = "add"
	TypeSetDefault = "setdefault"
	TypeConfig     = "config"
	TypeCancel     = "cancel"
)

const (
	ConflictStrategyRename    = "rename"
	ConflictStrategyAsk       = "ask"
	ConflictStrategyOverwrite = "overwrite"
	ConflictStrategySkip      = "skip"
)

func ConflictStrategyValues() []string {
	return []string{
		ConflictStrategyRename,
		ConflictStrategyAsk,
		ConflictStrategyOverwrite,
		ConflictStrategySkip,
	}
}

func ConflictStrategyDisplay(strategy, lang string) string {
	displays := map[string]map[string]string{
		ConflictStrategyRename: {
			"zh-CN": "始终重命名",
			"en":    "Always rename",
		},
		ConflictStrategyAsk: {
			"zh-CN": "每次询问",
			"en":    "Ask every time",
		},
		ConflictStrategyOverwrite: {
			"zh-CN": "始终覆盖",
			"en":    "Always overwrite",
		},
		ConflictStrategySkip: {
			"zh-CN": "始终跳过",
			"en":    "Always skip",
		},
	}
	if display, ok := displays[strategy]; ok {
		if str, ok := display[lang]; ok {
			return str
		}
		return display["en"]
	}
	return strategy
}

func IsConflictStrategy(strategy string) bool {
	for _, value := range ConflictStrategyValues() {
		if strategy == value {
			return true
		}
	}
	return false
}

// type TaskDataTGFiles struct {
// 	Files   []tfile.TGFileMessage
// 	AsBatch bool
// }

// type TaskDataTelegraph struct {
// 	Pics     []string
// 	PageNode *telegraph.Page
// }

// type TaskDataType interface {
// 	TaskDataTGFiles | TaskDataTelegraph
// }

type Add struct {
	// [TODO] maybe we should to spilit this into different types...
	TaskType         tasktype.TaskType
	SelectedStorName string
	DirID            uint
	SettedDir        bool
	SelectedDirPath  string
	ConflictStrategy string
	// tfiles
	Files   []tfile.TGFileMessage
	AsBatch bool
	// tphpics
	TphPageNode *telegraph.Page
	TphPics     []string
	TphDirPath  string // unescaped telegraph.Page.Path
	// parseditem
	ParsedItem *parser.Item
	// directlinks
	DirectLinks []string
	// aria2
	Aria2URIs []string
	// ytdlp
	YtdlpURLs  []string
	YtdlpFlags []string
	// transfer
	TransferSourceStorName string
	TransferSourcePath     string
	TransferFiles          []string // file paths relative to source storage
}

type SetDefaultStorage struct {
	StorageName string
	DirID       uint
}
