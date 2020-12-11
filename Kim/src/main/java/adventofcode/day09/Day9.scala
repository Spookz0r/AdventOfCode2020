package adventofcode.day09

import scala.collection.mutable
import scala.io.Source
import scala.util.control.Breaks.{break, breakable}

object Day9 {
  def main(args: Array[String]): Unit = {
    def solve1: (Int, Int) = {
      val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day09/input.txt"
      val lines = Source.fromFile(filename).getLines.toList

      val preambleLen = 25;
      val preamble = new Preamble(preambleLen)
      for (i <- 0 until preambleLen) preamble.add(lines(i).toInt)
      for (i <- preambleLen until lines.size) {
        val number = lines(i).toInt

        if (!hasSum(number)) {
          return (number, i)
        }

        preamble.add(number)
      }

      def hasSum(current: Int): Boolean = {
        for (i <- preamble.numbers) {
          for (j <- preamble.numbers) {
            if (i + j == current) {
              return true
            }
          }
        }
        false
      }

      (-1, -1)
    }

    val (invalidNumber, invalidIndex) = solve1
    println("invalid number " + invalidNumber + " invalid index " +invalidIndex)
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day09/input.txt"
    val lines = Source.fromFile(filename).getLines.toList
    for (i <- 0 until invalidIndex) {
      var sum = 0
      var min = Int.MaxValue
      var max = Int.MinValue
      breakable {
        for (j <- i until invalidIndex) {
          val number = lines(j).toInt
          sum += number
          min = Math.min(min, number)
          max = Math.max(max, number)
          if (sum == invalidNumber) {
            println(min + max)
            return
          } else if (sum > invalidNumber) {
            break
          }
        }
      }
    }
  }


  class Preamble(length: Int) {
    private val maxLen = length
    val numbers = new mutable.ArrayBuffer[Int]()

    def add(item: Int): Unit = {
      numbers += item
      if (numbers.length > maxLen) numbers.remove(0)
    }

  }
}
