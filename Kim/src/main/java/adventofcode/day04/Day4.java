package adventofcode.day04;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Day4 {
  public static void main(String[] args) throws FileNotFoundException {
//    solve(0);
    solve(1);
  }

  private static void solve(int problem) throws FileNotFoundException {
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/day04/input.txt"));
    List<String> passports = new ArrayList<>();
    String line = "";
    while (scanner.hasNextLine()) {
      String read = scanner.nextLine();
      if (read.isEmpty()) {
        passports.add(line);
        line = "";
      } else {
        line = line.isEmpty() ? read : line + " " + read;
      }
    }
    if (!line.isEmpty())
      passports.add(line);

    int validPassports = 0;
    for (String passport : passports) {
      Set<String> required = Stream.of("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid").collect(Collectors.toSet());
      for (String keyPair : passport.split(" ")) {
        if (problem != 0 && !isValueValid(keyPair.split(":")[0], keyPair.split(":")[1])) {
          break;
        }
        required.remove(keyPair.split(":")[0]);
      }
      if (required.isEmpty()) {
        validPassports++;
      }
      System.out.println(required.isEmpty() ? "VALID " + passport : "INVALID " + passport);
    }
    System.out.println(validPassports);
  }

  private static boolean isValueValid(String key, String value) {
    switch (key) {
      case "byr":
        if (value.length() == 4) return 1920 <= Integer.parseInt(value) && Integer.parseInt(value) <= 2002;
        else return false;
      case "iyr":
        if (value.length() == 4) return 2010 <= Integer.parseInt(value) && Integer.parseInt(value) <= 2020;
        else return false;
      case "eyr":
        if (value.length() == 4) return 2020 <= Integer.parseInt(value) && Integer.parseInt(value) <= 2030;
        else return false;
      case "hgt":
        if (value.endsWith("cm")) {
          int digit = Integer.parseInt(value.split("cm")[0]);
          return 150 <= digit && digit <= 193;
        } else if (value.endsWith("in")) {
          int digit = Integer.parseInt(value.split("in")[0]);
          return 59 <= digit && digit <= 76;
        } else {
          return false;
        }
      case "hcl":
        return value.length() == 7 && value.startsWith("#");
      case "ecl":
        return Stream.of("amb", "blu", "brn", "gry", "grn", "hzl", "oth").collect(Collectors.toSet()).contains(value);
      case "pid":
        return value.length() == 9 && value.matches("[0-9]{9}");
      case "cid":
        return true;
    }
    return false;
  }

}
