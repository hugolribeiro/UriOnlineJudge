import java.util.Scanner;

public class Main{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String values1 = sc.nextLine();
        String values2 = sc.nextLine();

        int numberUnits1 = Integer.parseInt(values1.split(" ")[1]);
        double priceProduct1 = Double.parseDouble(values1.split(" ")[2]);

        int numberUnits2 = Integer.parseInt(values2.split(" ")[1]);
        double priceProduct2 = Double.parseDouble(values2.split(" ")[2]);

        double total = (numberUnits1 * priceProduct1) + (numberUnits2 * priceProduct2);
        System.out.printf("VALOR A PAGAR: R$ %.2f\n", total);

        sc.close();
    }
}