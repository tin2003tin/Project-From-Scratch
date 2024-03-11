package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}
type Set struct {
	Text string
	Problem []string
	questions []Question
}
func (s *Set) ShuffleQuestions() {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(s.questions), func(i, j int) {
        s.questions[i], s.questions[j] = s.questions[j], s.questions[i]
    })
}

type Game struct {
	sets []Set
}

func main() {
	// Read the file
	file, err := os.Open("C:/Users/com/Documents/GitHub/Tcp/client/Obejct/unit1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Parse the file
	scanner := bufio.NewScanner(file)
	var questions []Question
	var problem []string;
	var game Game
	rand.Seed(time.Now().UnixNano())
	var readMode bool = false;
	var gameMode bool = true;
	fmt.Println(readMode,gameMode,problem)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "//") {
			s := Set{Text: line[2:]}
			game.sets = append(game.sets, s);
			continue
		}

		if strings.HasPrefix(line, "<<") {
			readMode = true
			questions = nil
			problem = nil;
			continue
		}

		if strings.HasPrefix(line, ">>") {
			readMode = false
			s := Set{questions: questions,Problem: problem}
			game.sets = append(game.sets, s)
			continue
		}
		if (readMode && strings.HasPrefix(line, "?")) {
			problem = append(problem, line[1:])
			continue
		}
		if (readMode) {
			qa := strings.Split(line, " ")
			questions = append(questions, Question{Text:qa[0],Answer: qa[1] })
		}
		
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	curr := 0;
	for {
		if (curr >= len(game.sets)) {
			break;
		}
		set := game.sets[curr];
		if (set.Text != "") {
			fmt.Println(set.Text)
		} else {
			for _, text := range set.Problem {
			fmt.Println(text)
			}
			
			q_index := 0
			set.ShuffleQuestions();
			for {
				if (q_index >= len(set.questions)) {
					break;
				}
				currentQuestion := set.questions[q_index];
				fmt.Println(q_index+1, currentQuestion.Text)
				reader := bufio.NewReader(os.Stdin)
				userInput, _ := reader.ReadString('\n')
				userInput = strings.TrimSpace(userInput)

				if strings.ToUpper(userInput) == strings.ToUpper(currentQuestion.Answer) {
					fmt.Println("Correct!")
				} else if userInput == "skip" {
					fmt.Println("Skipping...")
					break
				} else {
					fmt.Println("Incorrect! The correct answer is:", currentQuestion.Answer)
				}
				q_index++
			}
			fmt.Println();
		}
		curr++;
	}
}
