import java.math.BigDecimal;
import java.util.Scanner;
import java.math.RoundingMode;


public class Main {
    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        String sellerName = sc.nextLine();
        double fixedSalary = sc.nextDouble();
        double totalSale = sc.nextDouble();
        double total = fixedSalary + (totalSale * 0.15);
        
        System.out.println(String.format("TOTAL = R$ %.2f" , total));
        
        sc.close();
    }
}
