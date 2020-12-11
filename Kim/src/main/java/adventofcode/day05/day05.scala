package adventofcode.day05

import scala.collection.mutable.ArrayBuffer
import scala.io.Source

object day05 {
  def main(args: Array[String]): Unit = {
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day05/input.txt"
    var max = 0
    val list = ArrayBuffer[Int]()
    for (line <- Source.fromFile(filename).getLines) {
      var minr = 0
      var maxr = 127
      var minc = 0
      var maxc = 7
      line.foreach(c => {
        if (c == 'F') {
          maxr = minr + ((maxr-minr) / 2);
        } else if (c == 'B') {
          minr = minr + (maxr-minr) / 2 + 1
        } else if (c == 'L') {
          maxc = minc + ((maxc-minc) / 2)
        } else if (c == 'R') {
          minc = minc + ((maxc-minc)/2 + 1)
        }
      })

      max = math.max(max, minr * 8 + minc)
      list.addOne(minr * 8 + minc);
    }
    //println(max)
    list.sortInPlaceWith(_ < _)
    for (i <- 1 until list.size) {
      if (list(i) - list(i-1) > 1) {
        println("My seat is " + (list(i) - 1))
        return
      }
    }

  }
}
