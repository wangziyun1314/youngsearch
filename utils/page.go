package utils

type Page struct {
	TotalNumber   int         // 搜索出的总记录的格式
	PageNumber    int         // 一共有多少页
	CurrentPage   int         // 当前页
	PerPageNumber int         // 每个页面应该具有的记录数量
	PageContent   [][2]string // 页面当前的内容
}

func GetPages(msg [][2]string) []Page {
	total := len(msg)
	i := total / PageLimit
	pageNumber := 0
	if total%PageLimit == 0 {
		pageNumber = i
	} else {
		pageNumber = i + 1
	}
	pages := make([]Page, pageNumber)
	for i := 1; i <= pageNumber; i++ {
		content := make([][2]string, 0)
		for j := 0; j < PageLimit; j++ {
			if (i-1)*PageLimit+j < len(msg) {
				content = append(content, msg[(i-1)*PageLimit+j])
			}
		}
		pages[i-1] = Page{
			TotalNumber:   total,
			PageNumber:    pageNumber,
			CurrentPage:   i,
			PerPageNumber: PageLimit,
			PageContent:   content,
		}
	}
	return pages
}
