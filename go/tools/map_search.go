package main

import (
	dbm "GmTest/db_modules"
	"GmTest/mysql_module"
	"fmt"
	"gorm.io/gorm/logger"
	"reflect"
	"sort"
	"time"
)

// Point 表示二维点
type Point struct {
	X float64
	Y float64
}

// Quadrant 表示象限
type Quadrant int

const (
	NW Quadrant = iota
	NE
	SW
	SE
)

// QuadTreeNode 表示四叉树节点
type QuadTreeNode struct {
	Boundary struct {
		Min Point
		Max Point
	}
	Points   []Point
	Children [4]*QuadTreeNode
}

// QueryRange 范围查询
func QueryRange(node *QuadTreeNode, rangeBoundary struct{ Min, Max Point }) []Point {
	var result []Point
	queryRange(node, rangeBoundary, &result)
	return result
}

func queryRange(node *QuadTreeNode, rangeBoundary struct{ Min, Max Point }, result *[]Point) {
	if node == nil {
		return
	}

	// 如果节点边界不与查询范围相交，直接返回
	if node.Boundary.Min.X > rangeBoundary.Max.X || node.Boundary.Max.X < rangeBoundary.Min.X ||
		node.Boundary.Min.Y > rangeBoundary.Max.Y || node.Boundary.Max.Y < rangeBoundary.Min.Y {
		return
	}

	// 检查节点中的点是否在查询范围内
	for _, point := range node.Points {
		if point.X >= rangeBoundary.Min.X && point.X <= rangeBoundary.Max.X &&
			point.Y >= rangeBoundary.Min.Y && point.Y <= rangeBoundary.Max.Y {
			*result = append(*result, point)
		}
	}

	// 递归查询子节点
	for _, child := range node.Children {
		queryRange(child, rangeBoundary, result)
	}
}

// NewQuadTreeNode 创建新的四叉树节点
func NewQuadTreeNode(min, max Point) *QuadTreeNode {
	return &QuadTreeNode{
		Boundary: struct{ Min, Max Point }{Min: min, Max: max},
		Points:   make([]Point, 0),
		Children: [4]*QuadTreeNode{nil, nil, nil, nil},
	}
}

// Insert 插入点
func (n *QuadTreeNode) Insert(point Point) {
	if !n.inBoundary(point) {
		return
	}

	if len(n.Points) < 4 {
		// 按照坐标大小插入到合适的位置
		if len(n.Points) == 0 || point.X < n.Points[0].X || (point.X == n.Points[0].X && point.Y < n.Points[0].Y) {
			n.Points = append([]Point{point}, n.Points...)
		} else {
			n.Points = append(n.Points, point)
		}
	} else {
		if n.Children[NW] == nil {
			n.subdivide()
		}

		quadrant := n.getQuadrant(point)
		switch quadrant {
		case NW, NE, SW, SE:
			n.Children[quadrant].Insert(point)
		default:
			// Handle the case where a point lies exactly on the boundary
			n.Points = append(n.Points, point)
		}
	}
}

// inBoundary 判断点是否在节点的边界内
func (n *QuadTreeNode) inBoundary(point Point) bool {
	return point.X >= n.Boundary.Min.X &&
		point.X <= n.Boundary.Max.X &&
		point.Y >= n.Boundary.Min.Y &&
		point.Y <= n.Boundary.Max.Y
}

// getQuadrant 获取点所在象限
func (n *QuadTreeNode) getQuadrant(point Point) Quadrant {
	midX := (n.Boundary.Min.X + n.Boundary.Max.X) / 2.0
	midY := (n.Boundary.Min.Y + n.Boundary.Max.Y) / 2.0

	switch {
	case point.X < midX && point.Y < midY:
		return NW
	case point.X >= midX && point.Y < midY:
		return NE
	case point.X < midX && point.Y >= midY:
		return SW
	case point.X >= midX && point.Y >= midY:
		return SE
	default:
		return NW
	}
}

// subdivide 将节点分为四个子节点
func (n *QuadTreeNode) subdivide() {
	midX := (n.Boundary.Min.X + n.Boundary.Max.X) / 2.0
	midY := (n.Boundary.Min.Y + n.Boundary.Max.Y) / 2.0

	n.Children[NW] = NewQuadTreeNode(n.Boundary.Min, Point{X: midX, Y: midY})
	n.Children[NE] = NewQuadTreeNode(Point{X: midX, Y: n.Boundary.Min.Y}, Point{X: n.Boundary.Max.X, Y: midY})
	n.Children[SW] = NewQuadTreeNode(Point{X: n.Boundary.Min.X, Y: midY}, Point{X: midX, Y: n.Boundary.Max.Y})
	n.Children[SE] = NewQuadTreeNode(Point{X: midX, Y: midY}, n.Boundary.Max)

	// 将节点中的点重新分配给子节点
	for _, point := range n.Points {
		n.Children[n.getQuadrant(point)].Insert(point)
	}

	// 清空当前节点的点
	n.Points = nil
}

func main() {
	// 示例用法
	dns := "lzj:123456@tcp(127.0.0.1:3306)/dopai?charset=utf8mb4&parseTime=True&loc=Local"
	conn, db := mysql_module.NewConn(dns, logger.Error)
	defer conn.Close()
	maps := make([]dbm.DbYuanyuzhou, 0)
	_ = db.Table("yuanyuzhou").Find(&maps).Error
	var (
		err error
	)

	min := &Point{}
	max := &Point{}
	err = db.Table("yuanyuzhou").Select("min(x) as x,min(y) as y").Find(&min).Error
	err = db.Table("yuanyuzhou").Select("max(x) as x,max(y) as y").Find(&max).Error
	_ = err

	rootBoundary := struct{ Min, Max Point }{
		Min: Point{X: min.X, Y: min.Y},
		Max: Point{X: max.X, Y: max.Y},
	}
	rootNode := NewQuadTreeNode(rootBoundary.Min, rootBoundary.Max)
	for _, v := range maps {
		rootNode.Insert(Point{X: float64(v.X), Y: float64(v.Y)})
	}

	search1 := func() {
		queryRange := struct{ Min, Max Point }{Min: Point{X: -20, Y: -100}, Max: Point{X: 30, Y: 30}}
		checkPoint(QueryRange(rootNode, queryRange))
	}
	search2 := func() {
		result2 := make([]Point, 0)
		for _, v := range maps {
			if v.X >= -20 && v.X <= 30 && v.Y >= -100 && v.Y <= 30 {
				result2 = append(result2, Point{X: float64(v.X), Y: float64(v.Y)})
			}
		}
		checkPoint(result2)
	}
	useTime(search2)
	useTime(search1)
}

func checkPoint(result []Point) {
	sort.Slice(result, func(i, j int) bool {
		return result[i].X > result[j].X && result[i].Y > result[j].Y
	})
	fmt.Println(result[0])
	fmt.Println(result[5])
	fmt.Println(result[len(result)-1])
}

func useTime(f func()) {
	start := time.Now().Unix()
	for i := 1; i <= 1; i++ {
		f()
	}
	end := time.Now().Unix()
	t := reflect.TypeOf(f)
	fmt.Printf("f:%v use:%d\n", t.Name(), end-start)
}
