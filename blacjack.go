package main

import (
    "fmt"
    "math/rand"
)

// Card: Rank, Value
// Suits: 4 {13}
// Deck: 52 Cards
// Hand: Array Of Cards

//var RANKS = []string{"A", "J", "K", "Q", "10", "9", "8", "7",  "6", "5", "4", "3", "2", "1"}
var SUITS = []string{"SPADE", "DIAMOND", "CLUB", "HEART"}
var RANKS = map[string]int {"A": 11,"J":10, "K":10, "Q":10, "10":10, "9":9, "8":8, "7":7,  "6":6, "5":5, "4":4, "3":3, "2":2}

type Printer interface {
    Print()
}

type Card struct {
    suit, rank string
    value int
}

type Deck struct {
    cards []Card
}

type Hand struct {
    cards []Card
}

func (deck *Deck) init() {
    for _, suit := range SUITS {
        for rank, value := range RANKS {
            card := Card{suit, rank, value}
            deck.cards = append(deck.cards, card)
        }
    }
}

func (deck *Deck) shuffle() {
    for i := range deck.cards {
        j := rand.Intn(i+1)
        deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
    }
}

func (deck *Deck) drawCard() Card {
    card := deck.cards[0]
    deck.cards = append(deck.cards[:0], deck.cards[1:]...)
    return card
}

func (hand *Hand) hit(deck *Deck) {
    hand.cards = append(hand.cards, deck.drawCard())
    if hand.count() > 21 {
        fmt.Println("Player Busted!!! Dealer Wins!!!")
        return
    }
}

func (hand *Hand) Print() {
    fmt.Print("Printing Hand: ")
    for _, v := range hand.cards {
        fmt.Print(v, " ")
    }
    fmt.Println()
}

func (deck *Deck) Print() {
    fmt.Print("Printing Deck: ")
    for _, v := range deck.cards {
        fmt.Print(v, " ")
    }
    fmt.Println()
}

func (hand *Hand) count() int {
    sum := 0
    num_aces := 0
    for _, card := range hand.cards {
        if card.rank == "A" {
            num_aces += 1
        }
        sum += card.value
    }
    for i:=0; i<num_aces; i++ {
        if sum > 21 {
            sum -= 10
        }
    }
    return sum
}

func (hand *Hand) stand(dealer *Hand) {
    if dealer.count() - hand.count() == 0 {
        fmt.Println("Draw!!!")
    } else if hand.count() > 21 {
        fmt.Println("Player Busted!!! Dealer Wins.")
    } else if dealer.count() > 21 {
        fmt.Println("Dealer Busted!!! Player Wins.")
    } else {
        fmt.Println("Player Wins.")
    }
}

func main () {
    var choice string
    d := Deck{}
    d.init()
    d.shuffle()
    
    fmt.Println("Starting BlackJack: ")
    for {
        player := Hand{}
        dealer := Hand{}

        player.hit(&d)
        player.hit(&d)
        dealer.hit(&d)
        dealer.hit(&d)

        player.Print()
        dealer.Print()

        fmt.Println(player.count())
        fmt.Println(dealer.count())

        for {
            fmt.Print("Do you want to  HIT/STAND [H/S]? ")
            fmt.Scanln(&choice)
            if choice == "H" {
                player.hit(&d)
                dealer.hit(&d)
                player.Print()
                dealer.Print()

                fmt.Println(player.count())
                fmt.Println(dealer.count())

            } else {
                player.stand(&dealer)
                player.Print()
                dealer.Print()

                fmt.Println(player.count())
                fmt.Println(dealer.count())
                break
            }
        }
        fmt.Print("Do You Want To Play Again [Y/N]? ")
        fmt.Scanln(&choice)
        if choice == "N" {
            return
        }
    }
}

