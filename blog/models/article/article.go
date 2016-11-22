package article

import (
	m "blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*新增文章*/
// 分类有三种情况：
// 1.id为空,parentid为空，表示初始化分类
// 2.id为空，parentid不为空，表示新增次级分类
// 3.id不为空，parentid为空，表示不新增分类
func InsertArticle(article m.Article) int64 {
	m.GetRom().Begin()
	category := *article.Category
	var cid = category.Id
	if cid == 0 {
		cid, _ = m.GetRom().Insert(&category)
	}
	newArticle := &article
	newArticle.Category.Id = cid //添加外键方式
	id, err := m.GetRom().Insert(newArticle)
	commitChekErr(err)
	return id

}

func UpdateArticle(article m.Article) int64 {
	m.GetRom().Begin()
	category := *article.Category
	var cid = category.Id
	if cid == 0 {
		cid, _ = m.GetRom().Insert(&category)
	}
	newArticle := &article
	newArticle.Category.Id = cid //添加外键方式
	_, err := m.GetRom().Update(newArticle)
	commitChekErr(err)
	return article.Id
}

/*获取文章对象*/
func GetArticleById(typeStr string, id int64) (*m.Article, int64) {
	article := &m.Article{}
	if typeStr == "a" {
		//根据id查询文章对象
		// m.GetRom().QueryTable("t_article").Filter("Id", id).RelatedSel().One(article)
		id = QueryDataRowCountById(id)

	} else if typeStr == "n" {
		//根据行数查询对象，默认
	}

	m.GetRom().QueryTable(new(m.Article)).Limit(1, id-1).OrderBy("id").RelatedSel().One(article)
	return article, id
}

func QueryArticleByCategory(categoryId int64, num, pageNum, total int) ([]m.Article, []orm.Params) {
	var articles []m.Article
	var maps []orm.Params
	m.GetRom().Raw("SELECT id, title,time FROM t_article order by  time desc limit 0,?", num).QueryRows(&articles)
	m.GetRom().Raw("select a.id,a.title,a.time,a.summary,c.id category_id,c.name from t_article a inner join t_category c on c.id = a.category_id where c.id=? order by a.time desc limit ?,?", categoryId, total*(pageNum-1), total).Values(&maps)
	return articles, maps
}

func DelArticleById(id int64) bool {
	if _, err := m.GetRom().Delete(&m.Article{Id: id}); err == nil {
		return true
	}
	return false
}

func GetArticleByCategoryId(cid int64) int64 {
	cnt, _ := m.GetRom().QueryTable(new(m.Article)).Filter("category_id", cid).Count()
	return cnt
}

func QueryHomeData(num, pageNum, total int) ([]m.Article, []orm.Params) {
	var articles []m.Article
	var maps []orm.Params
	m.GetRom().Raw("SELECT id, title,time FROM t_article order by  time desc limit 0,?", num).QueryRows(&articles)
	m.GetRom().Raw("select a.id,a.title,a.time,a.summary,c.id category_id,c.name from t_article a inner join t_category c on c.id = a.category_id order by a.time desc limit ?,?", total*(pageNum-1), total).Values(&maps)
	return articles, maps
}

func QueryDataRowCountById(id int64) int64 {
	cnt, _ := m.GetRom().QueryTable(new(m.Article)).OrderBy("id").Filter("id__lte", id).Count()
	return cnt
}

func commitChekErr(err error) {
	if err == nil {
		m.GetRom().Commit()
		fmt.Println("提交。。。")
	} else {
		m.GetRom().Rollback()
		fmt.Println("回滚。。。", err)
	}
}
