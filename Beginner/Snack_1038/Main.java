import java.util.HashMap;
import java.util.Hashtable;
import java.util.Map;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        Map<String, Double> snacks = new HashMap<String, Double>();
        snacks.put("1", 4.0);
        snacks.put("2", 4.5);
        snacks.put("3", 5.0);
        snacks.put("4", 2.0);
        snacks.put("5", 1.5);

        
        String values = sc.nextLine();
        String[] splittedValues = values.split(" ");
        String code = splittedValues[0];
        double amount = Double.parseDouble(splittedValues[1]);
        Double total = snacks.get(code) * amount;
        
        System.out.printf("Total: R$ %.2f\n", total);
        sc.close();
    }
}
