package main 

import (
    "strings"
)

//Calculate the likelihood a string is any of a list of strings.
//Max 1000 strings for now to ensure that it runs quickly...
//...recommended to run in batches
func algo(w1 string, ws []string, wsl int) [1000]float64 {
    var reward = 0
    var penalty = 0
    w1m := map[string]int{}
    var output [1000]float64
    var w2 string
    num := 1000
    //if there are less than 1000, don't run 1000 times
    if(wsl<num) {
        num = wsl
    }
    //create a map of the first word's substrings for repeatability
    for x := 0; x < len(w1); x++ {
        for y := x+1; y < len(w1); y++ {
            w1m[w1[x:y]] += 1
        }
    }
    //iterate through the given words
    for wi:=0; wi<num; wi++ {
        //set up penalty, reward, and given word
        reward = 0
        penalty = 0
        w2 = ws[wi]
        //if the list word's substrings (exhaustive) exist within the map...
        for x := 0; x < len(w2); x++ {
            for y := x+1; y < len(w2); y++ {
                if(w1m[w2[x:y]] != 0) {
                    //...reward
                    reward += (y-x)*2
                } else {
                    //...penalize
                    penalty += 1
                }
             }
         }
         //if the first word's map's keys within the list word...
         for x, s := range w1m {
             if(strings.Contains(w2, x)) {
                 //..reward
                 reward += len(x)*2*s
             } else {
                 //...penalize
                 penalty += 1
             }
        }
        //Score is bound between 0 and 1
        output[wi] = float64(reward)/float64(reward+penalty)
    }
    return output
}

