import java.util.Scanner;

public class Main {
    
    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        int employeeNumber = sc.nextInt();
        int workedHours = sc.nextInt();
        double amountPerHour = sc.nextDouble();
        double salary = amountPerHour * workedHours;

        System.out.printf("NUMBER = %d\n", employeeNumber);
        System.out.printf("SALARY = U$ %.2f\n", salary);

        sc.close();

    }
}
