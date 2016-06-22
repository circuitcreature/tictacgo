package main
import(
	"fmt"
	"bytes"
	"strconv"
)

type game struct{
	set [9]string
	size int
}

func(g *game)runner(x int, i int)bool{
	if g.set[x] == g.set[x+i] && g.set[x+i] == g.set[x+i+i]{
		return true
	}
	return false
}

func(g *game)Validate( i int )bool{
	//only works if size is odd
	if g.runner((i / g.size)*g.size, 1){
		fmt.Println("Checking: 1",i)
		return true
	}
	if g.runner(i % g.size, g.size){
		fmt.Println("Checking: 2",i)
		return true
	}

	switch i{
		case g.size - 1: fallthrough
		case g.size * 2:
			if g.runner(g.size - 1, g.size - 1){
				fmt.Println("Checking: 3",i)
				return true
			}
		case 0: fallthrough
		case g.size*g.size:
			if g.runner(0, g.size + (g.size / 2)){
				fmt.Println("Checking: 4",i)
				return true
			}
		case (g.size*g.size)/2:
			if g.runner(g.size - 1, g.size - 1) || g.runner(0, g.size + (g.size / 2)) {
				fmt.Println("Checking: 5",i)
				//diagonals
				return true
			}
	}
	return false
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
	row := 1
	new := true
	buffer.WriteString("\n   ")
	for k := 0; k < g.size; k++{
		buffer.WriteString(strconv.Itoa(k+1)+" ")
	}
	buffer.WriteString("\n")
	for k,_ := range g.set{
		if new {
			buffer.WriteString(strconv.Itoa(row)+" ")
			new = false
			row++
		}
		buffer.WriteString("|")
		if g.set[k] == ""{
			out = "_"
		}else{
			out = g.set[k]
		}
		buffer.WriteString(out)

		if pi%g.size == 0{
			buffer.WriteString("|")
			buffer.WriteString("\n")
			new = true
		}
		pi++
	}
	return buffer.String()
}

func (g *game)GetInput()int{
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil || i > g.size || i < 0 {
		fmt.Print("Please enter a valid number.")
		return g.GetInput()
	}
	return i
}

func main(){
	g := game{set:[9]string{}, size: 3}
	winner := false
	player := true
	count := 0
	cur := "X"
	for false == winner{
		fmt.Println(g)

		if !player{
			cur = "O"
		}else {
			cur = "X"
		}
		fmt.Println("Player ",cur,":")

		fmt.Println("Select a Row:")
		r := g.GetInput()

		fmt.Println("Select a Column:")
		c := g.GetInput()
		i := (r - 1) * 3 + c - 1

		if !g.Set( i, cur ){
			fmt.Println("Position is already taken.")
			continue
		}

		count++
		if len(g.set) == count{
			break
		}
		winner = g.Validate(i)
		player = !player
	}

	fmt.Println(g)
	if winner {
		fmt.Println(cur, "is the winner.")
	}else{
		fmt.Println("Cats")
	}
}
