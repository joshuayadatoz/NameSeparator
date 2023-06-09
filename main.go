package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	Title      string
	FirstName  string
	Initial    string
	LastName   string
}