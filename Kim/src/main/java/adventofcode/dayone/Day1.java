package adventofcode.dayone;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Day1 {

  public static void main(String[] args) throws FileNotFoundException {
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/dayone/input.txt"));
    List<Long> numbers = new ArrayList<>();

    while (scanner.hasNextLong())
      numbers.add(scanner.nextLong());
    numbers.sort(Long::compareTo);

    for (int i = 0; i < numbers.size(); i++) {
      for (int j = 0; j < numbers.size(); j++) {
        for (int k = 0; k < numbers.size(); k++) {
          if (i == j || i == k || j == k) continue;
          if (numbers.get(i) + numbers.get(j) + numbers.get(k) > 2020) {
            break;
          } else if (numbers.get(i) + numbers.get(j) + numbers.get(k) == 2020) {
            System.out.println(numbers.get(i) * numbers.get(j) * numbers.get(k));
            return;
          }
        }
      }
    }
  }
}
