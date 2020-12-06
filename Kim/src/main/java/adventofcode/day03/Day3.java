package adventofcode.day03;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Day3 {
  public static void main(String[] args) throws FileNotFoundException {
    //solve1();

    List<List<Character>> map = readMap();
    long r1d1 = solve2(map, 1, 1);
    long r3d1 = solve2(map, 3, 1);
    long r5d1 = solve2(map, 5, 1);
    long r7d1 = solve2(map, 7, 1);
    long r1d2 = solve2(map, 1, 2);
    System.out.println(r1d1 * r3d1 * r5d1 * r7d1 * r1d2);
  }

  private static long solve2(List<List<Character>> map, int right, int down) {
    long trees = 0;
    int current = 0;
    for (int i = down; i < map.size(); i = i + down) {
      current = next(current, right, map.get(i).size());
      if (map.get(i).get(current) == '#') {
        trees++;
      }
    }
    return trees;
  }

  private static List<List<Character>> readMap() throws FileNotFoundException {
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/day03/input.txt"));
    List<List<Character>> map = new ArrayList<>();
    while (scanner.hasNextLine())
      map.add(scanner.nextLine().chars().mapToObj(i -> (char) i).collect(Collectors.toList()));
    return map;
  }

  private static int next(int current, int right, int rowSize) {
    for (int i = 0; i < right; i++) {
      int next = current + 1;
      current = next == rowSize ? 0 : next;
    }
    return current;
  }

  private static void solve1() throws FileNotFoundException {
    Scanner scanner = new Scanner(new File("src/main/java/adventofcode/day03/input.txt"));
    List<List<Character>> map = new ArrayList<>();
    while (scanner.hasNextLine())
      map.add(scanner.nextLine().chars().mapToObj(i -> (char) i).collect(Collectors.toList()));

    int trees = 0;
    int current = 0;
    int increment = 3;
    for (int i = 1; i < map.size(); i++) {
      for (int j = 0; j < increment; j++) {
        int next = current + 1;
        current = next == map.get(i).size() ? 0 : next;
      }
      if (map.get(i).get(current) == '#') {
        trees++;
      }
    }
    System.out.println(trees);
  }
}
