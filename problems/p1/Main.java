
/**
 * Ron Nathaniel
 * July 26 2020
 */
 
public class Main {

  public static int getTotal(int limit) {
    int total = 0;
    for (int i = 0; i < limit; ++i) {
      if (i % 3 == 0 || i % 5 == 0) {
        total += i;
      }
    }

    return total;
  }

  public static void main(String[] args) {
    int t = getTotal(1000);
    System.out.printf("%d\n", t);
  }

}
