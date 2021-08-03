import java.util.Scanner;

public class Main {

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String[] values = sc.nextLine().split(" ");
        double a = Double.parseDouble(values[0]);
        double b = Double.parseDouble(values[1]);
        double c = Double.parseDouble(values[2]);
        
        double triangleArea = a * c / 2;
        double circleArea = 3.14159 * Math.pow(c, 2);
        double trapeziumArea = (a + b) * c / 2;
        double squareArea = b * b;
        double rectangleArea = a * b;

        System.out.printf("TRIANGULO: %.3f\n", triangleArea);
        System.out.printf("CIRCULO: %.3f\n", circleArea);
        System.out.printf("TRAPEZIO: %.3f\n", trapeziumArea);
        System.out.printf("QUADRADO: %.3f\n", squareArea);
        System.out.printf("RETANGULO: %.3f\n", rectangleArea);

        sc.close();
    }
}
