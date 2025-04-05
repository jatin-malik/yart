package main

import (
	"fmt"
	"github.com/jatin-malik/yart/canvas"
	"github.com/jatin-malik/yart/color"
	"github.com/jatin-malik/yart/geometry"
)

type Projectile struct {
	position geometry.Point
	velocity geometry.Vector
}

type Environment struct {
	gravity geometry.Vector
	wind    geometry.Vector
}

// tick simulates one unit of time in projectile's motion and considering this simulation,
// updates position and velocity of the projectile.
func tick(projectile *Projectile, env *Environment) {
	projectile.position = projectile.position.Add(projectile.velocity)
	projectile.velocity = projectile.velocity.Add(env.gravity).Add(env.wind)
}

func main() {
	// Initialize the projectile

	initialPosition := geometry.NewPoint(0, 1, 0)
	intialVelocity, _ := geometry.NewVector(1, 1.8, 0).Normalize()

	projectile := Projectile{
		position: initialPosition,
		velocity: intialVelocity.Multiply(11.25),
	}

	// initialize the environment
	env := Environment{
		gravity: geometry.NewVector(0, -0.1, 0),
		wind:    geometry.NewVector(-0.01, 0, 0),
	}

	// Init the canvas
	canvas := canvas.New(900, 550)

	// Run tick until the projectile hits the ground

	totalTicks := 0
	maxHeight := 0.0

	h := canvas.GetHeight()
	projectileColor := color.New(1, 0, 0)

	// Simulate the motion
	for {
		canvas.WritePixel(int(projectile.position.GetX()), h-1-int(projectile.position.GetY()), projectileColor)
		tick(&projectile, &env)
		height := projectile.position.GetY() - initialPosition.GetY()
		maxHeight = max(height, maxHeight)
		totalTicks++
		if projectile.position.GetY() <= 0 {
			// projectile is on ground
			break
		}
	}
	canvas.WritePixel(int(projectile.position.GetX()), h-1, projectileColor)
	canvas.ToDisk("projectile.ppm")

	distance := projectile.position.GetX() - initialPosition.GetX()
	report(totalTicks, distance, maxHeight)

}

func report(totalTicks int, distance float64, height float64) {
	fmt.Println()
	fmt.Printf("The motion took %d ticks\n", totalTicks)
	fmt.Printf("Projectile covered %f distance horizontally\n", distance)
	fmt.Printf("Projectile reached %f maximum height during motion\n", height)
}
