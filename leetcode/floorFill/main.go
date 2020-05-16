package main

import (
	"fmt"
	"sync"
)

func main() {
	image := [][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2

	fmt.Println(floodFillC(image, sr, sc, newColor))
}

func floodFillC(image [][]int, sr int, sc int, newColor int) [][]int {
	targetCollor := image[sr][sc]
	if newColor == targetCollor {
		return image
	}
	image[sr][sc] = newColor

	numOfWorker := 5
	jobChan := make(chan []int, 4*numOfWorker)
	taskDone := make(chan int, numOfWorker)
	newTaskChan := make(chan []int, 4*numOfWorker)
	taskCountMux := sync.Mutex{}
	taskCount := 0

	// job dispatch
	go func() {
		for newTask := range newTaskChan {
			fmt.Println(newTask)
			jobChan <- newTask
		}
	}()

	// assign first job
	go func() {
		newTaskChan <- []int{sr, sc}
	}()

	// launch workers
	for i := 0; i < numOfWorker; i++ {
		go floodFillWorker(image, targetCollor, newColor, jobChan, taskDone, newTaskChan, &taskCountMux, &taskCount)
	}

	for range taskDone {
		opTaskCount(&taskCountMux, &taskCount, false)
		if taskCount == 0 {
			break
		}
	}

	close(jobChan)
	close(taskDone)
	close(newTaskChan)

	return image
}

func opTaskCount(taskCountMux *sync.Mutex, taskCount *int, isNew bool) {
	taskCountMux.Lock()
	if isNew {
		*taskCount++
	} else {
		*taskCount--
	}
	taskCountMux.Unlock()
}

func floodFillWorker(image [][]int, targetColor, newColor int, job <-chan []int, taskDone chan<- int, newTaskChan chan<- []int, taskCountMux *sync.Mutex, taskCount *int) {
	for pixel := range job {
		x, y := pixel[0], pixel[1]
		fmt.Println(x, y)
		if x > 0 && image[x-1][y] == targetColor {
			image[x-1][y] = newColor
			opTaskCount(taskCountMux, taskCount, true)
			newTaskChan <- []int{x - 1, y}
		}

		if y > 0 && image[x][y-1] == targetColor {
			image[x][y-1] = newColor
			opTaskCount(taskCountMux, taskCount, true)
			newTaskChan <- []int{x, y - 1}
		}

		if x < len(image)-1 && image[x+1][y] == targetColor {
			image[x+1][y] = newColor
			opTaskCount(taskCountMux, taskCount, true)
			newTaskChan <- []int{x + 1, y}
		}

		if y < len(image[0])-1 && image[x][y+1] == targetColor {
			image[x][y+1] = newColor
			opTaskCount(taskCountMux, taskCount, true)
			newTaskChan <- []int{x, y + 1}
		}
		taskDone <- -1
	}

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	targetCollor := image[sr][sc]
	if newColor == targetCollor {
		return image
	}
	image[sr][sc] = newColor
	for targetPixels := [][]int{[]int{sr, sc}}; len(targetPixels) != 0; {
		lenTargetPixels := len(targetPixels)
		for i := 0; i < lenTargetPixels; i++ {
			pixel := targetPixels[i]
			x, y := pixel[0], pixel[1]
			if x > 0 && image[x-1][y] == targetCollor {
				targetPixels = append(targetPixels, []int{x - 1, y})
				image[x-1][y] = newColor
			}

			if y > 0 && image[x][y-1] == targetCollor {
				targetPixels = append(targetPixels, []int{x, y - 1})
				image[x][y-1] = newColor
			}

			if x < len(image)-1 && image[x+1][y] == targetCollor {
				targetPixels = append(targetPixels, []int{x + 1, y})
				image[x+1][y] = newColor
			}

			if y < len(image[0])-1 && image[x][y+1] == targetCollor {
				targetPixels = append(targetPixels, []int{x, y + 1})
				image[x][y+1] = newColor
			}
		}
		targetPixels = targetPixels[lenTargetPixels:]
	}

	return image
}
