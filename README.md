# goMeSH

goMeSH is a small library for parsing MeSH data files available at [https://www.nlm.nih.gov/mesh/download_mesh.html](https://www.nlm.nih.gov/mesh/download_mesh.html)

The code is probably best as an an example of how to parse large xml files (using
using only the go standard library). The functions in XMLFileParsers are based on
the design in this article [Parsing Huge XML Files With Go](http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/).

## Example
An example of using one of the parser functions would be as follows
```go
func () {
  f, err := os.Open("testdata/pa2017.xml")
  if err != nil {
    panic(err)
  }

  pac, errc := gomesh.ParsePharmacologicalActionSet(f)
  defer close(errc)


  count := 0
  for {
    select {
    case pa := <-pac:
      doSomethingWithPharmacologicalAction(pa)
      count++
    case err := <-errc:
      if err != io.EOF {
        panic(err)
      }

      log.Printf("Pharmacological Actions parsed: %d", count)
      return
    }
  }
}()
```

## Notes
I (crhntr) do not use this library in any production code nor do I commit to maintaining it.
