import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);

        int a = sc.nextInt();
        int b = sc.nextInt();
        int c = sc.nextInt();
        int d = sc.nextInt();

        int diferanca = a * b - c * d;
        System.out.printf("DIFERENCA = %d\n", diferanca);

        sc.close();
    }
    
}
