package main
import(
	"fmt"
	"bytes"
)

func init(){
/*
	[0][1][2]
	[3][4][5]
	[6][7][8]

	[0][1][2][3]
	[4][5][6][7]
	[8][9][A][B]
	[C][D][E][F]

	0  0 0
	1  1 0
	2  2 0
	3  3 0

	4  0 1
	5  1 1
	6  2 1
	7  3 1

	8  0 2
	9  1 2
	10 2 2
	11 3 2

	12 0 3
	13 1 3
	14 2 3
	15 3 3

	[ 0][ 1][ 2][ 3][ 4]
	[ 5][ 6][ 7][ 8][ 9]
	[10][11][12][13][14]
	[15][16][17][18][19]
	[20][21][22][23][24]

	0 0 0
	1 1 0
	2 2 0
	3 3 0
	4 4 0

	5 0 1
	6 1 1
	7 2 1
	8 3 1
	9 4 1

	10 0 2
	11 1 2
	12 2 2
	13 3 2
	14 4 2

	15 0 3
	16 1 3
	17 2 3
	18 3 3
	19 4 3

	20 0 4
	21 1 4
	22 2 4
	23 3 4
	24 4 4

	0{
		by1->0
		by3->0
		by4->0
	}
	1{
		by1->0
		by3->x
	}
	2{
		by1->0
		by2->x
		by3->x
	}
	3{
		by1->x
		by3->0
	}
	4{
		by1->x-1
		by2->x-2
		by3->x-3
		by4->x-4
	}
	5{
		by1->x-2
		by3->x-3
	}
	6{
		by1->x
		by2->x-4
		by3->x-6
	}
	7{
		by1->x-1
		by3->x-6
	}
	8{
		by1->x-2
		by3->x-6
		by4->x-8
	}

	0->0
	1->1
	2->2
	3->0
	4->1
	5->2
	6->0
	7->1
	8->2


*/
}
type game struct{
	set [9]string
	size int
}

func(g *game)Check(i int)bool{

	x := i / g.size
	//aka by 1
	if g.set[x] == g.set[x+1] && g.set[x] == g.set[x+2]{
		//test horizonal
		return true
	}

	x = i % g.size
	//aka by 3
	if g.set[x] == g.set[x+g.size] && g.set[x] == g.set[x+g.size+g.size]{
		//test vertical
		fmt.Println("validate first")
		return true
	}

	//if x == x we be in the middle

	//aka by 2
	x = i
	if g.set[x] == g.set[x+2] && g.set[x] == g.set[x+4]{
		fmt.Println("validate by 2")
		//test horizonal
		return true
	}

	//aka by 4
	//for testing
	x = 0
	if g.set[x] == g.set[x+4] && g.set[x] == g.set[x+8]{
		//test horizonal
		return true
	}


	return false
}

func(g *game)Validate( i int )bool{
	if !g.Check( i ){
		fmt.Println("validate: ", i )
		//every function will validate by_three
		return false
	}


	return true
}

func(g *game)Set(i int, s string)bool{
	if g.set[i] != ""{
		return false
	}
	g.set[i] = s
	return true
}

func (g game) String() string{
	var buffer bytes.Buffer
	out := ""
	pi := 1
	buffer.WriteString("\n")
	for k,_ := range g.set{
		buffer.WriteString("|")
		if g.set[k] == ""{
			out = "_"
		}else{
			out = g.set[k]
		}
		buffer.WriteString(out)

		if pi%3 == 0{
			buffer.WriteString("|")
			buffer.WriteString("\n")
		}
		pi++
	}
	return buffer.String()
}

func main(){
	g := game{set:[9]string{"X","","O",
							"", "","X",
							"O","",""},
			size: 3}

	fmt.Println("Lets play a game!", g)

	x := 1
	g.Set(x, "X")
	fmt.Println("Lets play a game!", g)
	res := g.Validate(x)
	fmt.Println("validate is", res)
	x =4
	g.Set(x, "O")
	res = g.Validate(x)
	fmt.Println("Lets play a game!", g)
	fmt.Println("validate is", res)



}
