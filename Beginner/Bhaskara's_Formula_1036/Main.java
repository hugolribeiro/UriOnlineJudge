import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String[] inputtedValues = sc.nextLine().split(" ");
        double a = Double.parseDouble(inputtedValues[0]);
        double b = Double.parseDouble(inputtedValues[1]);
        double c = Double.parseDouble(inputtedValues[2]);

        double delta = Math.pow(b, 2) - (4 * a * c);
        if ((a == 0) || (delta < 0)) {
            System.out.println("Impossivel calcular");
        } else {
            double x1 = (-b + Math.sqrt(delta)) / (2 * a);
            double x2 = (-b - Math.sqrt(delta)) / (2 * a);
            System.out.printf("R1 = %.5f\n", x1);
            System.out.printf("R2 = %.5f\n", x2);
        }
        sc.close();
    }
}
