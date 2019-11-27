package gosnippets

import (
	"os"
	"os/exec"
	// "path"
	"bytes"
	"fmt"
	"strings"
	"bufio"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func ScreenCls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Watcher() {
	// _, filename, _, _ := runtime.Caller(1)
	// // path.Base(filename)
	// var watch = "go run" + path.Base(filename)
	// fmt.Print("->%s",watch)
	// // cmd := exec.Command("filewatcher", ".", "go run hello.go")
	// // cmd.Run()
	// // return filename
}

// return 0, and not 100 ,
// as 0 is the default value for an integer
func Defer() {

	aValue := new(int)

	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}

func RepeatString(st string, n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(st)
	}
	return buffer.String()
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func Prepend() {
	data := []string{"A", "B", "C", "D"}
	data = append([]string{"E"}, data...)
	fmt.Println(data)
	// [E A B C D]

	a := []int{2, 3, 4}
	b := 1
	fmt.Println(append([]int{b}, a...))
	// [1 2 3 4]
}

func Closure() int {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		a := 0
		if a > 0 {
			f()
		}
	}

	return 0
}

func TestEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// cannot convert i (type int) to type []byte
func TestConversion() {
	var i = 10
	b := []int{i}
	fmt.Println("%v", b)
}

func last_element_of_split() {
	s := "one.two.three"
	fmt.Println(s)
	fmt.Println(strings.Split(s, ".")[len(strings.Split(s, "."))-1])
}

func ZeroValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func ReadFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a,&b)
	res = a + b
	fmt.Println(res)
	
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseString2(s string) (result string) {
	for _, v  := range s {
		result = string(v) + result
	}
	return 
}

func ReverseWords(s string) string {
	// convert the string into an array of strings where each entry is a word
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func IsPalindrome(s string) string {
	mid := len(s) / 2
	last := len(s) - 1

	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "No " + s + " is not Palindrome"
		}
	}
	return "Yes " + s + " is Palindrome"
}

func RemoveDuplicatesFromSlice(s []string) []string {
      m := make(map[string]bool)
      for _, item := range s {
              if _, ok := m[item]; ok {
                      // duplicate item
                      //fmt.Println(item, "is a duplicate")
              } else {
                      m[item] = true
              }
      }

      var result []string
      for item, _ := range m {
              result = append(result, item)
      }
      return result
}

func ReplaceNumbersByZeroFromString () {
	str := "abc123465@ahsjf222"
    newStr := make([]rune, len(str))
    i, added := 0, false
    for _, r := range str {
        if r >= '0' && r <= '9' {
            if added {
                continue
            }
            added, newStr[i] = true, '0'
        } else {
            added, newStr[i] = false, r
        }
        i++
    }
    fmt.Println(string(newStr[:i]))
}

// Get back all environment variable
func OsEnviron() {
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
		fmt.Printf("KEY:%q\tVAL:%q\n", kv[0], kv[1])
	}
}

