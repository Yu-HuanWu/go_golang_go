package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("your have", fuel, "amount of fuel!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  return fuel
}

// Create the function greetPlanet() here


// Create the function cantFly() here


// Create the function flyToPlanet() here


func main() {
  // Test your functions!
  fuelGauge(42)
  // Create `planetChoice` and `fuel`
  
  // And then liftoff!
  
}