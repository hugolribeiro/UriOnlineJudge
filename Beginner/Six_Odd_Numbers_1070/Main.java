package  Beginner.Six_Odd_Numbers_1070;

import java.util.Scanner;

public class Main{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int number = sc.nextInt();
        int repeat_times = 6;

        if (number % 2 == 0) {
            number += 1;
        }
        for (int odd_number = number; odd_number < (number + (repeat_times * 2)); odd_number += 2) {
            System.out.println(odd_number);
        }
        sc.close();
    }
}