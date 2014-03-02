# GeoPoint

Go library to represents geographical point and formulaes to calculate distance to another geographical point. The formula supplied here is Haversine but you can also use your own formula.

This library is extremely simple to use:

    import (
      "fmt"
      "github.com/marcinwyszynski/geopoint"
    )
    
    func main() {
      dublin := geopoint.NewGeoPoint(53.347778, -6.259722)
      howth := geopoint.NewGeoPoint(53.386, -6.066)
      fmt.Printf("Author's daily commute is %.2f kilometres\n",
                 2 * howth.DistanceTo(dublin), geopoint.Haversine)
    }
    
For the full API please see [godoc.org](http://godoc.org/github.com/marcinwyszynski/geopoint).
