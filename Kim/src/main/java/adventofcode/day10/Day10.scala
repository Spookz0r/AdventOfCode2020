package adventofcode.day10

import scala.collection.mutable
import scala.io.Source

object Day10 {
  def main(args: Array[String]): Unit = {
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day10/input.txt"
    val adapters = Source.fromFile(filename).getLines.map(s => s.toInt).toList.sorted
    val max = adapters.max

    def nextAdapter(jolt: Int): (Int, Int) = {
      if (max == jolt) return (0,1)

      var diff1 = 0
      var diff3 = 0
      val next: Int = adapters.find(j => j > jolt && j <= jolt + 3).getOrElse(-1)
      if (next != -1) {
        val (count1, count3) = nextAdapter(next)
        diff1 += count1
        diff3 += count3
        if (next - jolt == 1) diff1 += 1
        if (next - jolt == 3) diff3 += 1
      }
      (diff1, diff3)
    }

    val (count1, count3): (Int, Int) = nextAdapter(0)
    print(count1 + "*" + count3 + "=" + (count1 * count3))
    println

    def nextAdapterExhaustive(jolt: Int, memo: mutable.HashMap[Int, Long]): Long = {
      if (max == jolt) return 1

      if (memo.contains(jolt)) return memo(jolt)

      var sum: Long = 0
      adapters.filter(j => j > jolt && j <= jolt + 3).foreach(next => {
        sum += nextAdapterExhaustive(next, memo)
      })

      memo(jolt) = sum
      memo(jolt)
    }

    val result = nextAdapterExhaustive(0, new mutable.HashMap[Int, Long]())
    println(result)


  }
}
