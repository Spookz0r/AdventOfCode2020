package adventofcode.day02;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day2 {
  public static void main(String[] args) throws FileNotFoundException {
    //solve1();
    solve2();
  }

  public static void solve2() throws FileNotFoundException {
    Pattern pattern = Pattern.compile("([0-9]*)-([0-9]*) ([a-z]): ([a-z]*)");
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/day02/input.txt"));
    int validPasswords = 0;
    String line;
    Matcher matcher;
    while (scanner.hasNextLine()) {
      line = scanner.nextLine();
      matcher = pattern.matcher(line);
      matcher.find();
      int min = Integer.parseInt(matcher.group(1)) - 1;
      int max = Integer.parseInt(matcher.group(2)) - 1;
      char letter = matcher.group(3).charAt(0);
      String password = matcher.group(4);

      if (password.charAt(min) == letter || password.charAt(max) == letter) {
        if (password.charAt(min) != password.charAt(max)) {
          validPasswords++;
        }
      }
    }
    System.out.println(validPasswords);
  }

  public static void solve1() throws FileNotFoundException {
    Pattern pattern = Pattern.compile("([0-9]*)-([0-9]*) ([a-z]): ([a-z]*)");
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/day02/input.txt"));
    int validPasswords = 0;
    String line;
    Matcher matcher;
    while (scanner.hasNextLine()) {
      line = scanner.nextLine();
      matcher = pattern.matcher(line);
      matcher.find();
      int min = Integer.parseInt(matcher.group(1));
      int max = Integer.parseInt(matcher.group(2));
      char letter = matcher.group(3).charAt(0);
      String password = matcher.group(4);

      long occurences = password.chars().mapToObj(i -> (char) i).filter(c -> c == letter).count();
      if (min <= occurences && occurences <= max) {
        validPasswords++;
      }
    }
    System.out.println(validPasswords);

  }
}
