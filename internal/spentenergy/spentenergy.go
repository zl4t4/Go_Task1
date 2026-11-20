package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func checkParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	return nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}
	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}
	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
