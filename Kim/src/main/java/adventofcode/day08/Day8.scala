package adventofcode.day08

import scala.collection.mutable.ArrayBuffer
import scala.io.Source
import scala.util.control.Breaks.{break, breakable}

object Day8 {
  def main(args: Array[String]): Unit = {
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day08/input.txt"
    val lines = Source.fromFile(filename).getLines.toList
    //solve1
    solve2

    def solve2: Unit = {
      val fromNop = attemptRepair("nop", "jmp")
      val fromJmp = attemptRepair("jmp", "nop")
      println(if (fromNop != -1) fromNop else fromJmp)
    }

    def attemptRepair(from: String, to: String): Int = {
      // change nop to jmp
      var changedIndex = -1
      breakable {
        while (true) {
          var instructions = new ArrayBuffer[String]
          var changed = false
          for (i <- lines.indices) {
            if (lines(i).startsWith(from) && changedIndex < i && !changed) {
              changedIndex = i
              changed = true
              instructions += lines(i).replace(from, to)
            } else
              instructions += lines(i)
          }
          if (!changed) break
          val result = runInstructions(instructions.toList)
          if (result != -1) return result
        }
      }
      -1
    }

    def runInstructions(instructions: List[String]): Int = {
      val visited = collection.mutable.Set[Int]()
      var acc = 0
      var i = 0
      breakable {
        while (i < instructions.size) {
          if (visited.contains(i)) {
            acc = -1
            break
          }

          visited += i
          instructions(i) match {
            case x if x.startsWith("acc") =>
              acc = handleAccumulator(x, acc)
              i += 1
            case x if x.startsWith("jmp") => i = handleJmp(x, i)
            case _ => i += 1
          }
        }
      }
      acc
    }

    def solve1 = {
      val visited = collection.mutable.Set[Int]()
      var acc = 0
      var i = 0
      breakable {
        while (i < lines.size) {
          if (visited.contains(i)) break

          visited += i
          lines(i) match {
            case x if x.startsWith("acc") => {
              acc = handleAccumulator(x, acc)
              i += 1
            }
            case x if x.startsWith("jmp") => i = handleJmp(x, i)
            case _ => i += 1
          }
        }
      }
      println(acc)
    }
  }

  def handleAccumulator(instruction: String, accumulator: Int): Int = {
    if (instruction.contains("+"))
      accumulator + instruction.split("\\+")(1).toInt else
      accumulator - instruction.split("-")(1).toInt
  }

  def handleJmp(instruction: String, index: Int): Int = {
    if (instruction.contains("+"))
      index + instruction.split("\\+")(1).toInt else
      index - instruction.split("-")(1).toInt
  }
}
