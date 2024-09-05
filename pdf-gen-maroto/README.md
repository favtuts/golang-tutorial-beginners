# Go long by generating PDFs in Golang with Maroto
* https://tuts.heomi.net/go-long-by-generating-pdfs-in-golang-with-maroto/

# Ensure latest Golang version

I am using [goenv](https://tuts.heomi.net/installing-multiple-version-of-golang-using-goenv/) to manage Go version in my machine.

```bash
# check all version which can be installed
$ goenv install -l

# install the latest version
$ goenv install 1.22.4

# make this version to be global
$ goenv global 1.22.4
$ goenv versions

# restart shell
$ exec $SHELL

# check go version
$ go version
go version go1.22.4 linux/amd64
```

# Initializing a new Go project

Locate the project location
```bash
$ mkdir pdf-gen-maroto
$ cd pdf-gen-maroto
```

Follow the following command:
```bash
$ go mod init github.com/USERNAME/PROJECT_NAME
```
In the command above, replace `github.com` with the domain where you store your files, such as `Bitbucket` or `Gitlab`. Additionally, substitute `USERNAME` with your username and `PROJECT_NAME` with the desired project name.

Run the command for init Go project
```bash
$ go mod init github.com/favtuts/maroto-pdf
```

To install Maroto as a dependency, run the following command:
```bash
$ go get github.com/johnfercher/maroto/v2@v2.1.2
```

For writing code, create a new file called main.go to the root of your project folder, and paste the following code into it:
```go
package main

import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
```

Now, run the command `go run main.go` from the terminal, and check if it prints `Hello, World!` in the terminal. If it does, it means you are ready to get started.


# Build the application

Let's start by creating a struct that defines the structure of the PDF:
```go
package main

type Company struct {
    Name         string
    Address      string
    LogoLocation string
}

type Ticket struct {
    ID                 int
    ShowName           string
    ShowTime           string
    Language           string
    ShowVenue          string
    SeatNumber         string
    Cost               float64
    Screen             string
    TicketCount        int
    ShowPosterLocation string
}

func main() {
    c := Company{
        Name:         "ShowBees Ticketing",
        Address:      "1234 Main St, City, State 12345",
        LogoLocation: "./logo.png",
    }

    t := Ticket{
        ID:                 1,
        ShowName:           "Planet of the Gophers: The War Begins",
        ShowTime:           "Sat 01/01/2022 7:00 PM",
        Language:           "English",
        ShowVenue:          "Gophedorium",
        SeatNumber:         "Platinum - A1, A2",
        Cost:               620.00,
        Screen:             "Screen 1",
        TicketCount:        2,
        ShowPosterLocation: "./poster.png",
    }
}
```

# Creating a header

The new function will take a parameter of type `Company` and will return a `core.Row`. The header is pretty simple. Let’s check out the code first:
```go
func getPageHeader(c Company) core.Row {
    return row.New(16).Add(
        image.NewFromFileCol(4, c.LogoLocation, props.Rect{
            Center:  false,
            Percent: 100,
        }),
        col.New(2),
        col.New(6).Add(
            text.New(c.Name, props.Text{
                Style: fontstyle.Bold,
                Size:  10,
            }),
            text.New(c.Address, props.Text{
                Top:  6,
                Size: 10,
            }),
        ),
    )
}
```

The header consists of three columns: an image column, an empty column, and a column containing text components.

Let’s now create a function called `getMaroto` which will be responsible for registering necessary components and returning a `core.Maroto` interface which wraps the basic methods of Maroto.

```go
func getMaroto(c Company, t Ticket) core.Maroto {
    cfg := config.NewBuilder().WithDimensions(120, 200).Build()

    mrt := maroto.New(cfg)

    err := mrt.RegisterHeader(getPageHeader(c))

    if err != nil {
        log.Println("Error registering header")
    }

    return mrt
}
```

At this point, we only have a header component. As we add more components to the PDF, this function will grow larger. The function takes two parameters: a company object and a ticket object.

Update the main function as shown below:
```go
func main() {
  // ...
    // ...

    m := getMaroto(c, t)

    document, err := m.Generate()

    filename := fmt.Sprintf("ticket-%d.pdf", t.ID)

    if err != nil {
        log.Println("Error generating PDF")
    }

    // Check if temp folder exists, if not create it
    if _, err := os.Stat("temp"); os.IsNotExist(err) {
        err = os.Mkdir("temp", 0755)
        if err != nil {
            log.Println("Error creating directory:", err)
        }
    }

    err = document.Save("temp/" + filename)
    if err != nil {
        log.Println("Unable to save file:", err)
    }
}
```

If you run the code using `go run main.go`, you’ll be able to see that a new folder called temp is created, and the folder contains a file called ticket-1.pdf.

# Fixed: Error loading workspace: err: exit status 1: stderr: go: updates to go.sum needed, disabled by -mod=readonly : packages.Load error
* https://stackoverflow.com/questions/67800641/error-loading-workspace-err-exit-status-1-stderr-go-updates-to-go-sum-neede

After so much searching, only this worked for me: (1) Disabling all extensions, (2) Closing VS Code and reopening it, (3) Enabling only Go extension


# Creating a body

Create a new function called `getShowDetails` which accepts a ticket struct and returns an array of `core.Row` interface. 
```go
func getShowDetails(t Ticket) []core.Row {
    rows := []core.Row{
        row.New(30).Add(
            image.NewFromFileCol(4, t.ShowPosterLocation, props.Rect{
                Center:  true,
                Percent: 100,
            }),
            col.New(8).Add(
                text.New(t.ShowName, props.Text{
                    Style: fontstyle.Bold,
                    Size:  10,
                }),
                text.New(t.Language, props.Text{
                    Top:   6,
                    Style: fontstyle.Normal,
                    Size:  8,
                    Color: &props.Color{Red: 95, Green: 95, Blue: 95},
                }),
                text.New(t.ShowTime, props.Text{
                    Top:   12,
                    Style: fontstyle.Bold,
                    Size:  10,
                }),
                text.New(t.ShowVenue, props.Text{
                    Top:   18,
                    Style: fontstyle.Normal,
                    Size:  8,
                    Color: &props.Color{Red: 95, Green: 95, Blue: 95},
                }),
            ),
        ),
        row.New(6),
        row.New(1).Add(
            line.NewCol(12, props.Line{
                Thickness:   0.2,
                Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
                SizePercent: 100,
                Style:       linestyle.Dashed,
            }),
        ),
        row.New(3),
        row.New(16).Add(
            col.New(2).Add(
                text.New(strconv.Itoa(t.TicketCount), props.Text{
                    Style: fontstyle.Bold,
                    Size:  24,
                    Align: align.Center,
                }),
                text.New("Tickets", props.Text{
                    Top:   12,
                    Style: fontstyle.Normal,
                    Size:  8,
                    Color: &props.Color{Red: 95, Green: 95, Blue: 95},
                    Align: align.Center,
                }),
            ),
            col.New(2),
            col.New(8).Add(
                text.New(t.Screen, props.Text{
                    Size:  8,
                    Color: &props.Color{Red: 95, Green: 95, Blue: 95},
                }),
                text.New(t.SeatNumber, props.Text{
                    Top:   6,
                    Style: fontstyle.Bold,
                    Size:  14,
                }),
            ),
        ),
        row.New(3),
        row.New(1).Add(
            line.NewCol(12, props.Line{
                Thickness:   0.2,
                Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
                SizePercent: 100,
                Style:       linestyle.Dashed,
            }),
        ),
        row.New(6),
        row.New(20).Add(
            code.NewQrCol(12,
                fmt.Sprintf("%v\n%v\n%v\n%v", t.ID, t.ShowName, t.ShowTime, t.ShowVenue),
                props.Rect{
                    Center:  true,
                    Percent: 100,
                },
            ),
        ),
        row.New(10).Add(
            col.New(12).Add(text.New(fmt.Sprintf("Booking ID: %v", t.ID), props.Text{
                Style: fontstyle.Normal,
                Size:  8,
                Align: align.Center,
                Top:   2,
            })),
        ),
        row.New(1).Add(
            line.NewCol(12, props.Line{
                Thickness:   0.2,
                Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
                SizePercent: 100,
                Style:       linestyle.Solid,
            }),
        ),
        row.New(3),
        row.New(10).Add(
            code.NewBarCol(12, strconv.Itoa(t.ID),
                props.Barcode{
                    Center:  true,
                    Percent: 100,
                },
            ),
        ),
    }

    return rows
}
```

# Creating a footer

The footer contains just a single text. The aim is to demonstrate how a footer can be registered in your PDF:

```go
func getPageFooter() core.Row {
    return row.New(2).Add(
        col.New(12).Add(
            text.New("Powered by ShowBees Ticketing System", props.Text{
                Style: fontstyle.Italic,
                Size:  8,
                Align: align.Center,
                Color: &props.Color{Red: 255, Green: 120, Blue: 218},
            }),
        ),
    )
}
```


Now, to register this footer and the `getShowDetails` function into the `getMaroto` function, update this accordingly:
```go
func getMaroto(c Company, t Ticket) core.Maroto {
    cfg := config.NewBuilder().WithDimensions(120, 200).Build()

    // ...
    // ...

    mrt.AddRow(6)

    mrt.AddRow(4, line.NewCol(12, props.Line{
        Thickness:   0.2,
        Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
        SizePercent: 100,
    }))

    mrt.AddRow(6)

    mrt.AddRows(getShowDetails(t)...)

    mrt.AddRow(8)

    err = mrt.RegisterFooter(getPageFooter())

    if err != nil {
        log.Println("Error registering footer")
    }

    return mrt
}
```

If you generate the PDF by running `go run main.go`, you should get a PDF that resembles the image shown earlier in the article.