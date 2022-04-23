package main

import (
  "fmt"
)

type Graph struct{
  m[][]int                                               //гениальная матрица смежности
  name []string                                          //массив названий узлов
  size int                                               //размер графа
}
func(g *Graph) addEdge(i int,j int, weight int ) {       //добавление ребра от i к j с весом weight
  g.m[i][j] = weight
}
func(g *Graph) addNode(name string){                     //добавление узла с названием name
  g.name=append(g.name,name)
  g.size++

  rofl:=make([]int,g.size-1)                             //создание среза с рядом
  g.m = append(g.m,rofl)                                 //добавляем к существующей матрице
  for i := range g.m {                                   //в цикле добаляем столбец с 0-ми для созданных рядов
    g.m[i] = append(g.m[i],0)
  }

}
func(g *Graph) PrintMatrix() {                            //печать матрицы
  fmt.Printf("\x1b[46m%s\x1b[0m", "   ") // )
  // fmt.Print("   ")
  g.PrintNames()
  fmt.Println()
  for i := 0; i < g.size; i++ {
    fmt.Print(g.name[i], "  ")
    for j := 0; j < g.size; j++ {
      fmt.Print(g.m[i][j], "  ")
    }
    fmt.Println()
  }
}
func(g *Graph) PrintNames() {                             //печать названий узлов
  for i := 0; i < g.size; i++ {
    fmt.Print(g.name[i], "  ")
  }
}
func(g *Graph) Upgrade(){                                 //функция добавления узлов и ребер графа
  g.addNode("a")
  g.addNode("b")
  g.addNode("c")
  g.addNode("d")
  g.addNode("e")
  g.addNode("f")
  g.addNode("g")

  g.addEdge(0,1,1)
  g.addEdge(0,2,2) // for "a"

  g.addEdge(1,3,3) // for "b"

  g.addEdge(2,3,4)
  g.addEdge(2,4,5)
  g.addEdge(2,5,6) // for "c"

  g.addEdge(5,4,7)
  g.addEdge(5,6,8)// for "f"

}
func(g *Graph) DFS(from string, to string) bool{          //Depth-first search(поиск в глубину) для пользователя
  fromIndex := -1
  toIndex := -1
  for i := 0; i < g.size; i++ {                           //поиск в массиве с названиями, названий указанных пользователем
    if g.name[i] == from {                                //если нашли-сохраняем индекс
      fromIndex = i
    }
    if g.name[i] == to {                                  //если нашли-сохраняем индекс
      toIndex = i
    }
  }

  if fromIndex != -1 && toIndex != -1 {                   //проверка на нахождение указанных узлов
    var b [] bool
    for i:=0;i<g.size;i++{                                //создание+инициализация среза, который отвечает за начальное*
      b=append(b,false)                            //*состояние посещенных узлов
    }
    return g.supportDFS(toIndex,fromIndex,b,from)         //вызов главного DFS
  } else {                                                //если не нашли указанные названия
    fmt.Println("NO SUCH NAME")
    return false
  }
}
func(g *Graph) supportDFS(to int, curr int,visited[] bool, pathOfExile string) bool {   //support(ГЛАВНЫЙ) DFS
  if curr == to {                                         //проверка на выход из рекурсии, дошли ли до нужного узла
    pathOfExile += "->" + g.name[curr]                    //добавление к пути последнего узла
    fmt.Println("Path Of Exile: ", knife(pathOfExile))//печать пути, через функцию, уберающую повторяющиеся узлы
    return true
  }

  visited[curr] = true                                    //обновление информации о состоянии посещения текущего узла

  for j := 0; j < g.size; j++ {                           //цикл, отвечающий за проход по всем ребрам текущего узла
    if g.m[curr][j] != 0 && !visited[j] {                 //проверка на существование ребра + не посещен ли связанный узел
      pathOfExile += "->" + g.name[curr]                  //обновление пути

      deeper := g.supportDFS(to, j, visited, pathOfExile) //рекурсия к узлу, который не был посещен + к нему есть ребро

      if deeper {                                         //если рекурсия вернула true, то путь существует и найден
        return true
      }
    }
  }
  return false                                            //нет пути /
}

func(g *Graph) dijkstrasAlgorithm() {                     //алгоритм Дейкстры
  cur := 0                                                //создание нулевого индекса узла
  var weight []int                                        //создание среза для веса
  var visited []bool                                      //создание среза посещений узлов
  for i := 0; i < g.size; i++ {                           //инициализация срезов:
    weight = append(weight, 1000000)                        //вес-максимально большой
    visited = append(visited, false)                        //посещения-не посещенно
  }
  weight[0] = 0                                           //от начала до начала, как не странно расстояние 0 /
  for cur != -1 {                                         //главный цикл пока не выйдем за счет посещения всех узлов
    for i := 0; i < g.size; i++ {                         //цикл просмотра ребер у текущего узла
      if g.m[cur][i] + weight[cur] < weight[i] && g.m[cur][i] != 0 { //нахождение минимального веса от текущего узла
        weight[i] = g.m[cur][i] + weight[cur]             //его обновление, усли нашлось меньше
      }
    }
    visited[cur] = true                                   //считаем текущий узел посещенным
    cur = -1                                              //сброс текущего индекса
    minWeight := 1000000                                  //нициализация МАКСИМАЛЬНОГО веса
    for i := 0; i < g.size; i++ {                         //цикл для прохода по срезу посещений и веса
      if !visited[i] && weight[i] < minWeight {           //проверка узлов на посещение и на достигаемость
        cur = i                                           //какой-то из узлов не прошел??? = бан, обновление текущего*
        minWeight = weight[i]                             //и о5 главный цикл по новой /
      }
    }
  }
  fmt.Print("Min weight from first added node to:\n\t\t\t\t\t\t\t\t\t") //забор
  g.PrintNames()
  fmt.Println("\n\t\t\t\t\t\t\t\t\t",weight)          //вывод
}

func knife(path string) string{                           //функция очищает путь от повторяющихся узлов
  prev:=string(path[0])
  last:=string(path[0])
    for i := 1; i < len(path); i++ {                      //проход посимвольно по строчке пути
      if string(path[i])!=prev{                           //сравнение с предшествующим узлом
        last+=string(path[i])
        if i%3==0{                                        //обновление предшествующего узла
          prev=string(path[i])
        }
      }else{                                              //если тот же узел, то скип стрелочку(->)
        i+=2
      }
    }
  return last                                             //возврат идеального пути
}

func main() {
  g:=Graph{
    m:    nil,
    name: nil,
    size: 0,
  }
  g.Upgrade()
  g.PrintMatrix()

  fmt.Println()

  fmt.Println(g.DFS("a","g"))

  g.dijkstrasAlgorithm()
  fmt.Println("\n\n\n\n\n\n")

}