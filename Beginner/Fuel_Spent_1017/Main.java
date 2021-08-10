import java.util.Scanner;

public class Main{
    public static void main(String[] args){
        final int spentFuel = 12;
        Scanner sc = new Scanner(System.in);
        double spentTime = sc.nextInt();
        double averageSpeed = sc.nextInt();

        double distance = spentTime * averageSpeed / spentFuel;
        System.out.printf("%.3f\n", distance);
        sc.close();
    }
}