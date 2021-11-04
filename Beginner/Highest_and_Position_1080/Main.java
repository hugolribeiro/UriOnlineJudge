import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int highestNumber;
        int numberIndex = 1;
        int actualNumber;
        highestNumber = sc.nextInt();
        for (int index = 2; index <= 100; index++) {
            actualNumber = sc.nextInt();
            if (actualNumber > highestNumber) {
                highestNumber = actualNumber;
                numberIndex = index;
            }
        }
        System.out.println(highestNumber);
        System.out.println(numberIndex);
        sc.close();
    }
}