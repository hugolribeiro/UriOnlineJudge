import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        double number = sc.nextDouble();
        String message;

        if (number >= 0 && number <= 25) {
            message = "Intervalo [0,25]";
        } else if (number > 25 && number <= 50) {
            message = "Intervalo (25,50]";
        } else if (number > 50 && number <= 75) {
            message = "Intervalo (50,75]";
        } else if (number > 75 && number <= 100) {
            message = "Intervalo (75,100]";
        } else {
            message = "Fora de intervalo";
        }
        System.out.println(message);
        sc.close();
    }
}