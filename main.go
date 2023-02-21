package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string:`json:"id"`
	isbn string:`json:"isbn"`
	title string:`json:"title"`
	Director *Director:`json:"director"`
}
