import java.io.IOException;
import java.util.Scanner;
import java.util.Locale;
 
/**
 * IMPORTANT: 
 *      O nome da classe deve ser "Main" para que a sua solução execute
 *      Class name must be "Main" for your solution to execute
 *      El nombre de la clase debe ser "Main" para que su solución ejecutar
 */
public class Main {

    private static double PI = 3.14159;
 
    public static void main(String[] args) throws IOException {
 
        /**
         * Escreva a sua solução aqui
         * Code your solution here
         * Escriba su solución aquí
         */
        Locale.setDefault(Locale.US);

        Scanner sc = new Scanner(System.in);
        double radius = sc.nextDouble();
        double area = radius * radius * PI;
        System.out.printf("A=%.4f\n", area);

        sc.close();
        
 
    }
 
}