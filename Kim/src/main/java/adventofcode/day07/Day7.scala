package adventofcode.day07

import scala.collection.mutable
import scala.collection.mutable.ArrayBuffer
import scala.io.Source

object Day7 {
  def main(args: Array[String]): Unit = {
    val filename = "/home/kng/git/cp/adventofcode/adventofcode/src/main/java/adventofcode/day07/input.txt"
    val containers = new mutable.HashMap[String, ArrayBuffer[String]]

    for (line <- Source.fromFile(filename).getLines) {
      val content = line.split(" ");
      val bags: ArrayBuffer[String] = ArrayBuffer()
      var current = ""
      for (i <- content.indices) {
        current += " " + content(i)
        if (current.contains("contain") || current.contains(",") || current.contains(".")) {
          var content = current
            .replace("contain", "")
            .replace(",", "")
            .replace(".", "")
            .trim
          bags += content

          if (current.contains("contain")) containers.put(bags(0), ArrayBuffer())
          else if (!content.equals("no other bags")) {
            if (content.contains("1")) content = content.replace("bag", "bags")
            containers(bags(0)) += content
          }

          current = ""
        }
      }
    }
    //solve1
    solve2

    def solve2: Unit = {
      println(count("shiny gold bags") - 1)
      def count(key: String): Int = {
        var bags = 1
        for (s <- containers(key)) {
          val num = s.split(" ")(0).toInt
          val name = s.split(" ").slice(1, s.length).mkString(" ")
          bags += num * count(name)
        }
        bags
      }
    }


    def solve1: Unit = {
      val bags = new mutable.HashSet[String]
      var toCheck = Iterable("shiny gold bag")
      while (toCheck.nonEmpty) {
        toCheck = digDeeper(toCheck)
        println(bags.size)
      }

      def digDeeper(toCheck: Iterable[String]): mutable.HashSet[String] = {
        val toCheckNext = new mutable.HashSet[String]
        for (elem <- toCheck) {
          println(elem)
          toCheckNext.addAll(containers.filter(e => e._2.exists(s => s.contains(elem))).keys)
        }
        bags.addAll(toCheckNext)
        toCheckNext
      }
    }
  }
}
