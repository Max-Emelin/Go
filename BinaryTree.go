package main

import (
	"fmt"
)

type Tree struct {
	root *TreeNode
}
type  TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (tr *Tree) Print() {
	tr.root.Print()
	fmt.Println()
}
func (tr *Tree) ChangeRoot(tn *TreeNode) {
	tr.root=tn
}
func (tr *Tree) RootExists() bool  {
	if (tr.root!=nil){                              // проверка на существование корня
		return true
	}
	return false
}
func (tr *Tree) AddNode(value int) {
	if (tr.RootExists()) {                                //если корень есть, то вызываем для узла
		tr.root.AddNode(value)
	} else {
		tr.ChangeRoot(&TreeNode{value: value})            //если нет корня, то создаем
	}
}
func (tr *Tree) FindByValue(value int) *TreeNode{
	return tr.root.FindByValue(value)
}
func (tr *Tree) Delete(value int) {
	if (tr.RootExists()) {                                    //если корень есть
		if (tr.root.value != value) {                         //и его значение не равно нужному, то вызываем для узла
			tr.root.Delete(value)
		} else {                                              //если значение коря нужное,то
			oldLeft := tr.root.left                           //создаем указатели на право-лево у корня
			oldRight := tr.root.right
			replace := new(TreeNode)                          //создаем указатель на узел
			if tr.root.right != nil {                         //если справа есть элементы, то ищем минимальный
				replace = tr.root.right.MinValue()
			} else {                                          //если нету, то на замену корню идет левый
				replace = tr.root.left
			}
			if replace != nil {    //если мы удаляем последний корень, то у него replace=nil(думаю нестот удалять такое)
				tr.Delete(replace.value)
			}
			tr.ChangeRoot(replace)                             //меняем корень с помощью функции на replace
			if tr.root==nil{                                   //если корень стал пустым,то выходим
				return
			}
			if oldRight == tr.root { //если правый указатель будет указывать на корень, не присваеваем
				tr.root.left = oldLeft
				return
			}
			if oldLeft==tr.root{     //если левый указатель будет указывать на корень, не присваеваем
				return
			}
			tr.root.left = oldLeft   //если элемент не пренадлежал root.left или root.right, ему нужно их присвоить
			tr.root.right = oldRight
		}
	} else {
		fmt.Println("Nothing to delete!")
	}
}


func (tn *TreeNode) AddNode(value int) {
	if (value > tn.value) {										//если значение больше смотрим вправо
		if (tn.right != nil) {									//если справа есть узел, то рекурсия на него
			tn.right.AddNode(value)
		} else {                                                //если узла нет, то присваем новый
			tn.right = &TreeNode{value: value}
		}
	}else {
		if (value < tn.value) {                                 //если значение меньше смотрим влево
			if (tn.left != nil) { 								//если слева есть узел, то рекурсия на него
				tn.left.AddNode(value)
			} else {                                            //если узла нет, то присваем новый
				tn.left = &TreeNode{value: value}
			}
		}else{                                                  //если значение уже существует, то выводим сообщение
			fmt.Println("This node value already exists!")
		}
	}
}
func (tn *TreeNode) FindByValue(value int) *TreeNode {
	if(tn!=nil) {                                               //если переданный узел не nil, то идет сравнение
		if (value > tn.value) {                                 //если значение > то рекурсия на правый узел
			return tn.right.FindByValue(value)
		} else {                                                //если не > то или < или =
			if (value < tn.value) {                             //если значение < то рекурсия на левый узел
				return 	tn.left.FindByValue(value)
			} else{                                             //если не > и не < то нашли нужное значение
				return tn
			}
		}
	}else {                                                     //если переданный узел nil, то нету узла с данным значением
		fmt.Println("Node with value", value, "doesn't exist!")
		return nil
	}
}
func (tn *TreeNode) Print() {
	if tn == nil {                                              //если переданный узел nil то выводить нечего
		return
	}
	tn.left.Print()                                             //идет рекурсивная печать левых элементов
	if tn.left!=nil &&tn.right!=nil {                           //проверка можно ли обратиться к левому и правому эл
		fmt.Print(tn.value, "(", tn.left.value, " ", tn.right.value, ")\t\t")
	}else if tn.left!=nil{                                      //если есть только левый
		fmt.Print(tn.value, "(", tn.left.value, " nil)\t\t")
	}else if tn.right!=nil{                                     //если есть только правый
		fmt.Print(tn.value, "(nil ",tn.right.value,")\t\t")
	}else{                                                      //если нет никакого
		fmt.Println(tn.value, "(nil nil)\t\t")
	}
	tn.right.Print()                                            //идет рекурсивная печать правых элементов
}
func (tn *TreeNode) MinValue() *TreeNode {
	if tn!=nil {                                                //если переданный узел не пустой
		if tn.left != nil {                                     //если левый узел не пустой
			return tn.left.MinValue()                           //рекурсия влево
		} else{                                                 //если левый пустой, то переданный узел-наименьший
			return tn
		}
	}else{                                                      //если пустой,то получается мы передали указатель на nil
		return nil
	}
}
func (tn *TreeNode) Delete(value int) {
	if value > tn.value {                                            //если значение > текущего
		if tn.right != nil {                                         //если справа есть элемент
			if value > tn.right.value || value < tn.right.value {    //и правый элемент не равен значению, то рекурсия
				tn.right.Delete(value)
			} else {                                          //если правый элемент-нужный,то идут проверки как удолять
				if tn.right.right == nil && tn.right.left == nil {   //если у него нет указателей, то мы удаляем
					tn.right=nil
				}else if tn.right.right==nil{                 //если есть только правый, то указываем сразу на него
					tn.right=tn.right.left
				}else if tn.right.left==nil{                  //если есть только левый, то указываем сразу на него
					tn.right=tn.right.right
				}else {                                       //если есть оба, то сохраняем указатели текущего,
					oldLeft:=tn.right.left                    //находим минимальное справа значение, сохраняем его,
					oldRight:=tn.right.right                  //удаляем указатель на элемент с минимальным значением,
					replace:=tn.right.right.MinValue()        //ставим минимального на место удаляемого,
					tn.Delete(replace.value)                  //даем указатели удаляемого минимальному
					tn.right=replace
					tn.right.left=oldLeft
					tn.right.right=oldRight
				}
			}
		}else{                                                //если справа нету, значит и нужного нету, тк он больше
			fmt.Println("Nothing to delete!")
		}
	}else {
		if tn.left!=nil{
			if value > tn.left.value||value<tn.left.value{
				tn.left.Delete(value)
			}else{
				if tn.left.left==nil&&tn.left.right==nil{
					tn.left=nil
				}else if tn.left.right==nil{
					tn.left=tn.left.left
				}else if tn.left.left==nil{
					tn.left= tn.left.right
				}else{
					oldLeft:=tn.left.left
					oldRight:=tn.left.right
					replace:=tn.left.right.MinValue()
					tn.Delete(replace.value)
					tn.left=replace
					tn.left.left=oldLeft
					tn.left.right=oldRight
				}
			}
		}else{
			fmt.Println("Nothing to delete!")
		}
	}

}




func main() {
	hed:=new(TreeNode)
	hed.value=20

	zxc:=new(Tree)
	zxc.root=hed

	zxc.AddNode(50)
	zxc.AddNode(10)
	zxc.AddNode(11)
	zxc.AddNode(15)
	zxc.AddNode(13)
	zxc.AddNode(12)

	fmt.Println("Root:",zxc.root.value)
	zxc.Print()

	zxc.Delete(50)
	zxc.Print()

	fmt.Println(zxc.root)




}