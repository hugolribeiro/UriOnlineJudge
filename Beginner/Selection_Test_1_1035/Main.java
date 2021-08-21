import java.util.Scanner;

public class Main{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String[] input = sc.nextLine().split(" ");
        int size = input.length;
        int[] numbers = new int[size];
        for (int i = 0; i < size; i ++){
            numbers[i] = Integer.parseInt(input[i]);
        }
        String message;

        if (Tests(numbers[0], numbers[1], numbers[2], numbers[3])){
            message = "Valores aceitos";
        } else {
            message = "Valores nao aceitos";
        }
        System.out.println(message);

        sc.close();
    }

    public static boolean Tests(int a, int b, int c, int d){
        if (!(b > c && d > a)) {
            return false;
        }
        if (!(c+d > a+b)) {
            return false;
        }
        if (!(c > 0 && d > 0 && a%2 == 0)) {
            return false;
        }
        return true;
    }
}