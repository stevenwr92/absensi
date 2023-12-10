package handler

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stevenwr92/absensi/models"
	"github.com/stevenwr92/absensi/utils"
)

type Response struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func ClockIn(c *fiber.Ctx) error {
	// Client Ip dari Request
	// ip := c.IP()
	ip := "202.80.216.43"

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := int(claims["Id"].(float64))

	today := time.Now().Format("2006-01-02")
	var existingAttendance models.Attendance

	if err := utils.DB.
		Where("user_id = ? AND DATE(clock_in) = ?", userID, today).
		First(&existingAttendance).Error; err == nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "You have already clocked in for today")
	}

	response, err := utils.HitGeoApi(ip)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	responses := Response{}

	if err := json.Unmarshal(response.Body(), &responses); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse response"})
	}

	clockInTime := time.Now()

	attendance := models.Attendance{
		UserId:           userID,
		ClockIn:          clockInTime,
		ClockInIpAddress: ip,
		ClockInLongitude: responses.Longitude,
		ClockInLatitude:  responses.Latitude,
		CreatedAt:        clockInTime,
	}

	if err := utils.DB.Create(&attendance).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"Message": "Success Clock in"})
}

func ClockOut(c *fiber.Ctx) error {
	// Client Ip dari Request
	// ip := c.IP()

	ip := "202.80.216.43"

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := int(claims["Id"].(float64))

	today := time.Now().Format("2006-01-02")
	var existingAttendance models.Attendance

	if err := utils.DB.
		Where("user_id = ? AND DATE(clock_in) = ?", userID, today).
		First(&existingAttendance).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "You haven't clocked in for today")
	}

	if existingAttendance.ClockOut != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "You have already clocked out for today")
	}

	response, err := utils.HitGeoApi(ip)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	responses := Response{}

	if err := json.Unmarshal(response.Body(), &responses); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse response"})
	}

	clockOut := time.Now()

	existingAttendance.ClockOut = &clockOut
	existingAttendance.ClockOutIpAddress = ip
	existingAttendance.ClockOutLongitude = responses.Longitude
	existingAttendance.ClockOutLatitude = responses.Latitude

	if err := utils.DB.Save(&existingAttendance).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"Message": "Success Clock out"})

}

func GetAttendance(c *fiber.Ctx) error {

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := int(claims["Id"].(float64))

	attendances := []models.Attendance{}
	if err := utils.DB.Where("user_id = ?", userID).Find(&attendances).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(attendances)
}
