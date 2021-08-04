import java.util.Scanner;

public class Main {

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int totalDistance = sc.nextInt();
        double spentFuel = sc.nextDouble();
        double averageConsumption = totalDistance / spentFuel;
        System.out.printf("%.3f km/l\n", averageConsumption);
        sc.close();
    }
    
}
