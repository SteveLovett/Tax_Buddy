package main

type F8824 struct {
	// Form 8824 is used to report like-kind exchanges

	// Headers
	nameOnReturn string
	iDNumber string
	nameOnReturn2 string
	iDNumber2 string

	// Part I vars
	realProperty bool
	likeKindPropGivenUp string
	likeKindPropReceived string
	dateLikeKindPropAcquired string
	dateLikeKindPropTransferred string
	dateLikeKindPropReceivedIdentified string
	dateLikeKindPropActuallyReceived string
	relatedPartyTransaction bool

	// Part II vars
	nameOfRelatedParty string
	relationshipToTaxpayer string
	relatedPartyTIN string
	addressOfRelatedParty string
	relatedParty2YearDisposition bool
	taxpayer2YearDisposition bool
	dispositionDueToDeath bool
	dispositionInvoluntaryConversion bool
	establishNoTaxAvoidanceMotivation bool

	// Part III vars
	multiAssetExchange bool
	fMVOtherPropertyGivenUp int
	adjustedBasisOtherPropertyGivenUp int
	gainOrLossOtherPropertyGivenUp int
	propUsedAsAHome bool
	bootReceived int
	fMVLikeKindPropReceived int
	totalReceived int
	adjustedBasisLikeKindPropGaveUp int
	realizedGainOrLoss int
	lesserOfBootOrRealizedGainOrLoss int
	ordinaryIncomeUnderRecapture int
	lesserOfBootOrRealizedMinusRecapture int
	recognizedGain int
	deferredGainOrLoss int
	basisOfLikeKindPropReceived int

}

func (f8824 *F8824) NameOnReturn() string {