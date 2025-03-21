package entsoe

type Parameter string

const (
	ParameterDocumentType                                             = "documentType"
	ParameterDocStatus                                                = "docStatus"
	ParameterProcessType                                              = "processType"
	ParameterBusinessType                                             = "businessType"
	ParameterPsrType                                                  = "psrType"
	ParameterTypeMarketAgreementType                                  = "type_MarketAgreement.type"
	ParameterContractMarketAgreementType                              = "contract_MarketAgreement.Type"
	ParameterAuctionType                                              = "auction.Type"
	ParameterAuctionCategory                                          = "auction.Category"
	ParameterClassificationSequenceAttributeInstanceComponentPosition = "classificationSequence_AttributeInstanceComponent.Position"
	ParameterOutBiddingZoneDomain                                     = "outBiddingZone_Domain"
	ParameterBiddingZoneDomain                                        = "biddingZone_Domain"
	ParameterControlAreaDomain                                        = "controlArea_Domain"
	ParameterInDomain                                                 = "in_Domain"
	ParameterOutDomain                                                = "out_Domain"
	ParameterAcquiringDomain                                          = "acquiring_Domain"
	ParameterConnectingDomain                                         = "connecting_Domain"
	ParameterRegisteredResource                                       = "registeredResource"
	ParameterTimeInterval                                             = "TimeInterval"
	ParameterPeriodStart                                              = "periodStart"
	ParameterPeriodEnd                                                = "periodEnd"
	ParameterTimeIntervalUpdate                                       = "TimeIntervalUpdate"
	ParameterPeriodStartUpdate                                        = "PeriodStartUpdate"
	ParameterPeriodEndUpdate                                          = "PeriodEndUpdate"
	ParameterSecurityToken                                            = "securityToken"
)

type ContractMarketAgreementType string

const (
	ContractMarketAgreementTypeDaily    ContractMarketAgreementType = "A01"
	ContractMarketAgreementTypeWeekly   ContractMarketAgreementType = "A02"
	ContractMarketAgreementTypeMonthly  ContractMarketAgreementType = "A03"
	ContractMarketAgreementTypeYearly   ContractMarketAgreementType = "A04"
	ContractMarketAgreementTypeTotal    ContractMarketAgreementType = "A05"
	ContractMarketAgreementTypeLongTerm ContractMarketAgreementType = "A06"
	ContractMarketAgreementTypeIntraday ContractMarketAgreementType = "A07"
	ContractMarketAgreementTypeHourly   ContractMarketAgreementType = "A13"
)

type AuctionType string

const (
	AuctionTypeImplicit AuctionType = "A01"
	AuctionTypeExplicit AuctionType = "A02"
)

type AuctionCategory string

const (
	AuctionCategoryBase    AuctionCategory = "A01"
	AuctionCategoryPeak    AuctionCategory = "A02"
	AuctionCategoryOffPeak AuctionCategory = "A03"
	AuctionCategoryHourly  AuctionCategory = "A04"
)

type PsrType string

const (
	PsrTypeMixed                      PsrType = "A03"
	PsrTypeGeneration                 PsrType = "A04"
	PsrTypeLoad                       PsrType = "A05"
	PsrTypeBiomass                    PsrType = "B01"
	PsrTypeFossilBrownCoalLignite     PsrType = "B02"
	PsrTypeFossilCoalDerivedGas       PsrType = "B03"
	PsrTypeFossilGas                  PsrType = "B04"
	PsrTypeFossilHardCoal             PsrType = "B05"
	PsrTypeFossilOil                  PsrType = "B06"
	PsrTypeFossilOilShale             PsrType = "B07"
	PsrTypeFossilPeat                 PsrType = "B08"
	PsrTypeGeothermal                 PsrType = "B09"
	PsrTypeHydroPumpedStorage         PsrType = "B10"
	PsrTypeHydroRunOfRiverAndPoundage PsrType = "B11"
	PsrTypeHydroWaterReservoir        PsrType = "B12"
	PsrTypeMarine                     PsrType = "B13"
	PsrTypeNuclear                    PsrType = "B14"
	PsrTypeOtherRenewable             PsrType = "B15"
	PsrTypeSolar                      PsrType = "B16"
	PsrTypeWaste                      PsrType = "B17"
	PsrTypeWindOffshore               PsrType = "B18"
	PsrTypeWindOnshore                PsrType = "B19"
	PsrTypeOther                      PsrType = "B20"
	PsrTypeACLink                     PsrType = "B21"
	PsrTypeDCLink                     PsrType = "B22"
	PsrTypeSubstation                 PsrType = "B23"
	PsrTypeTransformer                PsrType = "B24"
)

// AllPsrTypes contains all PsrType constants
var AllPsrTypes = []PsrType{
	PsrTypeMixed,
	PsrTypeGeneration,
	PsrTypeLoad,
	PsrTypeBiomass,
	PsrTypeFossilBrownCoalLignite,
	PsrTypeFossilCoalDerivedGas,
	PsrTypeFossilGas,
	PsrTypeFossilHardCoal,
	PsrTypeFossilOil,
	PsrTypeFossilOilShale,
	PsrTypeFossilPeat,
	PsrTypeGeothermal,
	PsrTypeHydroPumpedStorage,
	PsrTypeHydroRunOfRiverAndPoundage,
	PsrTypeHydroWaterReservoir,
	PsrTypeMarine,
	PsrTypeNuclear,
	PsrTypeOtherRenewable,
	PsrTypeSolar,
	PsrTypeWaste,
	PsrTypeWindOffshore,
	PsrTypeWindOnshore,
	PsrTypeOther,
	PsrTypeACLink,
	PsrTypeDCLink,
	PsrTypeSubstation,
	PsrTypeTransformer,
}

type BusinessType string

const (
	BusinessTypeGeneralCapacityInformation           BusinessType = "A25"
	BusinessTypeAlreadyAllocatedCapacity             BusinessType = "A29"
	BusinessTypeRequestedCapacity                    BusinessType = "A43"
	BusinessTypeSystemOperatorRedispatching          BusinessType = "A46"
	BusinessTypePlannedMaintenance                   BusinessType = "A53"
	BusinessTypeUnplannedOutage                      BusinessType = "A54"
	BusinessTypeInternalRedispatch                   BusinessType = "A85"
	BusinessTypeFrequencyContainmentReserve          BusinessType = "A95"
	BusinessTypeAutomaticFrequencyRestorationReserve BusinessType = "A96"
	BusinessTypeManualFrequencyRestorationReserve    BusinessType = "A97"
	BusinessTypeReplacementReserve                   BusinessType = "A98"
	BusinessTypeInterconnectorNetworkEvolution       BusinessType = "B01"
	BusinessTypeInterconnectorNetworkDismantling     BusinessType = "B02"
	BusinessTypeCounterTrade                         BusinessType = "B03"
	BusinessTypeCongestionCosts                      BusinessType = "B04"
	BusinessTypeCapacityAllocated                    BusinessType = "B05"
	BusinessTypeAuctionRevenue                       BusinessType = "B07"
	BusinessTypeTotalNominatedCapacity               BusinessType = "B08"
	BusinessTypeNetPosition                          BusinessType = "B09"
	BusinessTypeCongestionIncome                     BusinessType = "B10"
	BusinessTypeProductionUnit                       BusinessType = "B11"
	BusinessTypeAreaControlError                     BusinessType = "B33"
	BusinessTypeProcuredCapacity                     BusinessType = "B95"
	BusinessTypeSharedBalancingReserveCapacity       BusinessType = "C22"
	BusinessTypeShareOfReserveCapacity               BusinessType = "C23"
	BusinessTypeActualReserveCapacity                BusinessType = "C24"
)

type ProcessType string

const (
	ProcessTypeDayAhead                             ProcessType = "A01"
	ProcessTypeIntraDayIncremental                  ProcessType = "A02"
	ProcessTypeRealised                             ProcessType = "A16"
	ProcessTypeIntradayTotal                        ProcessType = "A18"
	ProcessTypeWeekAhead                            ProcessType = "A31"
	ProcessTypeMonthAhead                           ProcessType = "A32"
	ProcessTypeYearAhead                            ProcessType = "A33"
	ProcessTypeSynchronisationProcess               ProcessType = "A39"
	ProcessTypeIntradayProcess                      ProcessType = "A40"
	ProcessTypeReplacementReserve                   ProcessType = "A46"
	ProcessTypeManualFrequencyRestorationReserve    ProcessType = "A47"
	ProcessTypeAutomaticFrequencyRestorationReserve ProcessType = "A51"
	ProcessTypeFrequencyContainmentReserve          ProcessType = "A52"
	ProcessTypeFrequencyRestorationReserve          ProcessType = "A56"
)

type DocStatus string

const (
	DocStatusIntermediate DocStatus = "A01"
	DocStatusFinal        DocStatus = "A02"
	DocStatusActive       DocStatus = "A05"
	DocStatusCancelled    DocStatus = "A09"
	DocStatusEstimated    DocStatus = "X01"
)

type DocumentType string

const (
	DocumentTypeFinalisedSchedule                        DocumentType = "A09"
	DocumentTypeAggregatedEnergyDataReport               DocumentType = "A11"
	DocumentTypeAcquiringSystemOperatorReserveSchedule   DocumentType = "A15"
	DocumentTypeBidDocument                              DocumentType = "A24"
	DocumentTypeAllocationResultDocument                 DocumentType = "A25"
	DocumentTypeCapacityDocument                         DocumentType = "A26"
	DocumentTypeAgreedCapacity                           DocumentType = "A31"
	DocumentTypeReserveAllocationResultDocument          DocumentType = "A38"
	DocumentTypePriceDocument                            DocumentType = "A44"
	DocumentTypeEstimatedNetTransferCapacity             DocumentType = "A61"
	DocumentTypeRedispatchNotice                         DocumentType = "A63"
	DocumentTypeSystemTotalLoad                          DocumentType = "A65"
	DocumentTypeInstalledGenerationPerType               DocumentType = "A68"
	DocumentTypeWindAndSolarForecast                     DocumentType = "A69"
	DocumentTypeLoadForecastMargin                       DocumentType = "A70"
	DocumentTypeGenerationForecast                       DocumentType = "A71"
	DocumentTypeReservoirFillingInformation              DocumentType = "A72"
	DocumentTypeActualGeneration                         DocumentType = "A73"
	DocumentTypeWindAndSolarGeneration                   DocumentType = "A74"
	DocumentTypeActualGenerationPerType                  DocumentType = "A75"
	DocumentTypeLoadUnavailability                       DocumentType = "A76"
	DocumentTypeProductionUnavailability                 DocumentType = "A77"
	DocumentTypeTransmissionUnavailability               DocumentType = "A78"
	DocumentTypeOffshoreGridInfrastructureUnavailability DocumentType = "A79"
	DocumentTypeGenerationUnavailability                 DocumentType = "A80"
	DocumentTypeContractedReserves                       DocumentType = "A81"
	DocumentTypeAcceptedOffers                           DocumentType = "A82"
	DocumentTypeActivatedBalancingQuantities             DocumentType = "A83"
	DocumentTypeActivatedBalancingPrices                 DocumentType = "A84"
	DocumentTypeImbalancePrices                          DocumentType = "A85"
	DocumentTypeImbalanceVolume                          DocumentType = "A86"
	DocumentTypeFinancialSituation                       DocumentType = "A87"
	DocumentTypeCrossBorderBalancing                     DocumentType = "A88"
	DocumentTypeContractedReservePrices                  DocumentType = "A89"
	DocumentTypeInterconnectionNetworkExpansion          DocumentType = "A90"
	DocumentTypeCounterTradeNotice                       DocumentType = "A91"
	DocumentTypeCongestionCosts                          DocumentType = "A92"
	DocumentTypeDcLinkCapacity                           DocumentType = "A93"
	DocumentTypeNonEuAllocations                         DocumentType = "A94"
	DocumentTypeConfigurationDocument                    DocumentType = "A95"
	DocumentTypeFlowBasedAllocations                     DocumentType = "B11"
)

type AreaType string

const (
	BZN AreaType = "Bidding Zone"
	BZA AreaType = "Bidding Zone Aggregation"
	CTA AreaType = "Control Area"
	MBA AreaType = "Market Balance Area"
	IBA AreaType = "Imbalance Area"
	IPA AreaType = "Imbalance Price Area"
	LFA AreaType = "Load Frequency Control Area"
	LFB AreaType = "Load Frequency Control Block"
	REG AreaType = "Region"
	SCA AreaType = "Scheduling Area"
	SNA AreaType = "Synchronous Area"
)

type DomainType = string

const (
	DomainNIR               DomainType = "10Y1001A1001A016"
	DomainEE                DomainType = "10Y1001A1001A39I"
	DomainSE1               DomainType = "10Y1001A1001A44P"
	DomainSE2               DomainType = "10Y1001A1001A45N"
	DomainSE3               DomainType = "10Y1001A1001A46L"
	DomainSE4               DomainType = "10Y1001A1001A47J"
	DomainNO5               DomainType = "10Y1001A1001A48H"
	DomainRussianArea       DomainType = "10Y1001A1001A49F"
	DomainRUKGD             DomainType = "10Y1001A1001A50U"
	DomainBelarusArea       DomainType = "10Y1001A1001A51S"
	DomainIESEM             DomainType = "10Y1001A1001A59C"
	DomainDEATLU            DomainType = "10Y1001A1001A63L"
	DomainNO1A              DomainType = "10Y1001A1001A64J"
	DomainDK                DomainType = "10Y1001A1001A65H"
	DomainITGR              DomainType = "10Y1001A1001A66F"
	DomainITNorthSI         DomainType = "10Y1001A1001A67D"
	DomainITNorthCH         DomainType = "10Y1001A1001A68B"
	DomainITBrindisi        DomainType = "10Y1001A1001A699"
	DomainITZCentreNorth    DomainType = "10Y1001A1001A70O"
	DomainITCentreSouth     DomainType = "10Y1001A1001A71M"
	DomainITZFoggia         DomainType = "10Y1001A1001A72K"
	DomainITNorth           DomainType = "10Y1001A1001A73I"
	DomainITZSardinia       DomainType = "10Y1001A1001A74G"
	DomainITSicily          DomainType = "10Y1001A1001A75E"
	DomainITZPriolo         DomainType = "10Y1001A1001A76C"
	DomainITRossano         DomainType = "10Y1001A1001A77A"
	DomainITZSouth          DomainType = "10Y1001A1001A788"
	DomainCADenmark         DomainType = "10Y1001A1001A796"
	DomainITNorthAT         DomainType = "10Y1001A1001A80L"
	DomainITNorthFR         DomainType = "10Y1001A1001A81J"
	DomainDELU              DomainType = "10Y1001A1001A82H"
	DomainDE                DomainType = "10Y1001A1001A83F"
	DomainITMACRZONENORTH   DomainType = "10Y1001A1001A84D"
	DomainITMACRZONESOUTH   DomainType = "10Y1001A1001A85B"
	DomainUADobTPP          DomainType = "10Y1001A1001A869"
	DomainITMalta           DomainType = "10Y1001A1001A877"
	DomainITSACOAC          DomainType = "10Y1001A1001A885"
	DomainITSACODC          DomainType = "10Y1001A1001A893"
	DomainNordic            DomainType = "10Y1001A1001A91G"
	DomainUK                DomainType = "10Y1001A1001A92E"
	DomainMT                DomainType = "10Y1001A1001A93C"
	DomainMD                DomainType = "10Y1001A1001A990"
	DomainAM                DomainType = "10Y1001A1001B004"
	DomainGE                DomainType = "10Y1001A1001B012"
	DomainAZ                DomainType = "10Y1001A1001B05V"
	DomainUA                DomainType = "10Y1001C--00003F"
	DomainUAIPS             DomainType = "10Y1001C--000182"
	DomainCZDESKLTSE4       DomainType = "10Y1001C--00038X"
	DomainCORE              DomainType = "10Y1001C--00059P"
	DomainAFRR              DomainType = "10Y1001C--00090V"
	DomainSWE               DomainType = "10Y1001C--00095L"
	DomainITCalabria        DomainType = "10Y1001C--00096J"
	DomainGBIFA             DomainType = "10Y1001C--00098F"
	DomainXK                DomainType = "10Y1001C--00100H"
	DomainIN                DomainType = "10Y1001C--00119X"
	DomainNO2A              DomainType = "10Y1001C--001219"
	DomainITALYNORTH        DomainType = "10Y1001C--00137V"
	DomainGRIT              DomainType = "10Y1001C--00138T"
	DomainAL                DomainType = "10YAL-KESH-----5"
	DomainAT                DomainType = "10YAT-APG------L"
	DomainBA                DomainType = "10YBA-JPCC-----D"
	DomainBE                DomainType = "10YBE----------2"
	DomainBG                DomainType = "10YCA-BULGARIA-R"
	DomainDEDK1LU           DomainType = "10YCB-GERMANY--8"
	DomainCBRSMKME          DomainType = "10YCB-JIEL-----9"
	DomainCBPL              DomainType = "10YCB-POLAND---Z"
	DomainCBSIHRBA          DomainType = "10YCB-SI-HR-BA-3"
	DomainCH                DomainType = "10YCH-SWISSGRIDZ"
	DomainME                DomainType = "10YCS-CG-TSO---S"
	DomainRS                DomainType = "10YCS-SERBIATSOV"
	DomainCY                DomainType = "10YCY-1001A0003J"
	DomainCZ                DomainType = "10YCZ-CEPS-----N"
	DomainDETransnetBW      DomainType = "10YDE-ENBW-----N"
	DomainDETenneTGER       DomainType = "10YDE-EON------1"
	DomainDEAmprion         DomainType = "10YDE-RWENET---I"
	DomainDE50Hertz         DomainType = "10YDE-VE-------2"
	DomainDK1A              DomainType = "10YDK-1-------AA"
	DomainDK1               DomainType = "10YDK-1--------W" 
	DomainDK2               DomainType = "10YDK-2--------M"
	DomainPLCZ              DomainType = "10YDOM-1001A082L"
	DomainCZDESK            DomainType = "10YDOM-CZ-DE-SKK"
	DomainLTSE4             DomainType = "10YDOM-PL-SE-LT2"
	DomainCWE               DomainType = "10YDOM-REGION-1V"
	DomainES                DomainType = "10YES-REE------0"
	DomainContinentalEurope DomainType = "10YEU-CONT-SYNC0"
	DomainFI                DomainType = "10YFI-1--------U"
	DomainFR                DomainType = "10YFR-RTE------C"
	DomainGB                DomainType = "10YGB----------A"
	DomainGR                DomainType = "10YGR-HTSO-----Y"
	DomainHR                DomainType = "10YHR-HEP------M"
	DomainHU                DomainType = "10YHU-MAVIR----U"
	DomainIE                DomainType = "10YIE-1001A00010"
	DomainIT                DomainType = "10YIT-GRTN-----B"
	DomainLT                DomainType = "10YLT-1001A0008Q"
	DomainLU                DomainType = "10YLU-CEGEDEL-NQ"
	DomainLV                DomainType = "10YLV-1001A00074"
	DomainMK                DomainType = "10YMK-MEPSO----8"
	DomainNL                DomainType = "10YNL----------L"
	DomainNO                DomainType = "10YNO-0--------C"
	DomainNO1               DomainType = "10YNO-1--------2"
	DomainNO2               DomainType = "10YNO-2--------T"
	DomainNO3               DomainType = "10YNO-3--------J"
	DomainNO4               DomainType = "10YNO-4--------9"
	DomainPL                DomainType = "10YPL-AREA-----S"
	DomainPT                DomainType = "10YPT-REN------W"
	DomainRO                DomainType = "10YRO-TEL------P"
	DomainSE                DomainType = "10YSE-1--------K"
	DomainSI                DomainType = "10YSI-ELES-----O"
	DomainSK                DomainType = "10YSK-SEPS-----K"
	DomainTR                DomainType = "10YTR-TEIAS----W"
	DomainUABEI             DomainType = "10YUA-WEPS-----0"
	DomainGBElecLink        DomainType = "11Y0-0000-0265-K"
	DomainGBIFA2            DomainType = "17Y0000009369493"
	DomainDK1NO1            DomainType = "46Y000000000007M"
	DomainNO2NSL            DomainType = "50Y0JVU59B4JWQCU"
	DomainBY                DomainType = "BY"
	DomainRU                DomainType = "RU"
	DomainIS                DomainType = "IS"
)

type ResolutionType string

const (
	ResolutionQuarter  ResolutionType = "PT15M"
	ResolutionHalfHour ResolutionType = "PT30M"
	ResolutionHour     ResolutionType = "PT60M"
	ResolutionDay      ResolutionType = "P1D"
	ResolutionWeek     ResolutionType = "P7D"
	ResolutionYear     ResolutionType = "P1Y"
)

