package adventofcode.day06

import scala.collection.mutable
import scala.io.Source

object Day6 {
  def main(args: Array[String]): Unit = {
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day06/input.txt"
    var result = 0
    var persons = 0
    var answers = mutable.Map[Char, Int]()
    for (line <- Source.fromFile(filename).getLines) {
      if (!line.isEmpty) {
        persons += 1
        line.foreach(c => {
          if (answers.contains(c)) answers.put(c, answers(c) + 1);
          else if (!answers.contains(c)) answers.put(c, 1)
        })
      }
      else {
        result += answers.values.count(_ == persons)
        answers = mutable.Map[Char, Int]()
        persons = 0
      }
    }
    result += answers.values.count(_ == persons)

    println(result)
  }
}
