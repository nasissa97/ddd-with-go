package patterns

import (
	"context"
	"errors"
	"fmt"

	"github.com/nasissa97/ddd-with-go/fundamentals"
)

type accountKey = int

const accountCtxKey = accountKey(1)

type BookingDomainService interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type BookingAppService struct {
	bookingRepo          BookingRepository
	bookingDomainService BookingDomainService
	emailService         EmailSender
}

func NewBookingAppService(
	bookingRepo BookingRepository,
	bookingDomainService BookingDomainService,
	emailInfraService EmailSender,
) *BookingAppService {
	return &BookingAppService{
		bookingRepo:          bookingRepo,
		bookingDomainService: bookingDomainService,
		emailService:         emailInfraService,
	}
}

func (b *BookingAppService) CreateBooking(ctx context.Context, booking Booking) error {
	u, ok := ctx.Value(accountCtxKey).(*fundamentals.Customer)
	if !ok {
		return errors.New("invalid customer")
	}

	if u.UserID() != booking.userID.String() {
		return errors.New("cannot create booking for other users")
	}

	if err := b.bookingDomainService.CreateBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not create booking: %w", err)
	}

	if err := b.bookingRepo.SaveBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not save booking: %w", err)
	}

	err := b.emailService.SendEmail(ctx, "me@example.com", "Test", "Practicing using Infra Service within App Service")
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
