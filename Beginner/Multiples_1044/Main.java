import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String values = sc.nextLine();
        String[] splittedValues = values.split(" ");
        int a = Integer.parseInt(splittedValues[0]);
        int b = Integer.parseInt(splittedValues[1]);

        if (a % b == 0 || b % a == 0) {
            System.out.println("Sao Multiplos");
        }
        else {
            System.out.println("Nao sao Multiplos");
        }
        sc.close();
    }
}
