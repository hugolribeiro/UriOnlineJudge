package Beginner.Interval2_1072;

import java.util.Scanner;

public class Main {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int amountNumbers = sc.nextInt();
        int amountIn = 0;
        int amountOut = 0;
        int actualNumber;
        for (int i = 0; i < amountNumbers; i ++){
            actualNumber = sc.nextInt();
            if (actualNumber < 10 || actualNumber > 20) {
                amountOut += 1;
                continue;
            }
            amountIn += 1;
        }
        System.out.printf("%d in\n", amountIn);
        System.out.printf("%d out\n", amountOut);
        sc.close();
    }
}
