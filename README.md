# Ambatustamp
# ![](assets/logo.png)
**Simple PDF stamping tool built with Go**
*It's so simple, it makes you want to stamp!*

## What is Ambatustamp?
Ambatustamp is a so-called "simple" PDF stamping tool built with Go. It claims to make document stamping easy, but don't get your hopes up.

## How do I use it?

### 1. Create a configuration struct
```go
stampConfig := ambatustamp.StampConfig{
   Size:     80,                  // in pixels
   LogoPath: "/path/to/logo.png", // or any other path
   Content:  uuid.New().String(), // or any other string
   Position: "bl",                // bl (bottom left), br (bottom right), tl (top left), tr (top right)
   Xoffset:  0,                   // -25 for padded stamp
   Yoffset:  0,                   // 25 for padded stamp
}

metadataConfig := ambatustamp.MetadataConfig{
   Title:   "Test PDF",             // or any other string
   Author:  "Ambatustamp",          // or any other string
   Subject: "Testing Ambatustamp",  // or any other string
}
```
### 2. Create an Ambatustamp instance
```go
amb := ambatustamp.NewAmbatustamp()
```
### 3. Load a PDF file
```go
err := amb.Load("/path/to/your/file.pdf")
```
### 3. If your PDF is password protected, provide the password
```go
err := amb.Decrypt("YOUR_PASSWORD")
```
### 4. Stamp the PDF
```go
err := amb.Stamp(&stampConfig)
```
### 5. Add metadata to the PDF
```go
err := amb.Metadata(&metadataConfig)
```
### 6. Save the PDF
```go
err := amb.Save("/path/to/your/file.pdf")
```

## How does it work? I know you're not a smart ass, so it must be simple.
Ambatustamp is nothing more than a lazy wrapper around another open-source project ([pdfcpu](https://github.com/pdfcpu/pdfcpu)). It's like putting a tuxedo on a couch potato.


