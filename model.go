/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package adv

import "time"

type ErrorMessage struct {
	Value string `json:"message"`
}

type Portfolio struct {
	Name    string `json:"name"`
	Uuid    string `json:"uuid"`
	Type    string `json:"type"`
	Deleted bool   `json:"deleted"`
}

type CfmFuturesPosition struct {
	ProductId         string `json:"product_id"`
	ExpirationTime    string `json:"expiration_time"`
	Side              string `json:"side"`
	NumberOfContracts string `json:"number_of_contracts"`
	CurrentPrice      string `json:"current_price"`
	AvgEntryPrice     string `json:"avg_entry_price"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
	DailyRealizedPnl  string `json:"daily_realized_pnl"`
}

type Account struct {
	Uuid              string `json:"uuid"`
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	AvailableBalance  Amount `json:"available_balance"`
	Default           bool   `json:"default"`
	Active            bool   `json:"active"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	Type              string `json:"type"`
	Ready             bool   `json:"ready"`
	Hold              Amount `json:"hold"`
	RetailPortfolioId string `json:"retail_portfolio_id"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type PriceBook struct {
	ProductId string  `json:"product_id"`
	Bids      []Level `json:"bids"`
	Asks      []Level `json:"asks"`
	Time      string  `json:"time"`
}

type Level struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

type ProductsResponse struct {
	Products    Product `json:"products"`
	NumProducts int     `json:"num_products"`
}

type Product struct {
	ProductId                 string               `json:"product_id"`
	Price                     string               `json:"price"`
	PricePercentageChange24h  string               `json:"price_percentage_change_24h"`
	Volume24h                 string               `json:"volume_24h"`
	VolumePercentageChange24h string               `json:"volume_percentage_change_24h"`
	BaseIncrement             string               `json:"base_increment"`
	QuoteIncrement            string               `json:"quote_increment"`
	QuoteMinSize              string               `json:"quote_min_size"`
	QuoteMaxSize              string               `json:"quote_max_size"`
	BaseMinSize               string               `json:"base_min_size"`
	BaseMaxSize               string               `json:"base_max_size"`
	BaseName                  string               `json:"base_name"`
	QuoteName                 string               `json:"quote_name"`
	Watched                   bool                 `json:"watched"`
	IsDisabled                bool                 `json:"is_disabled"`
	New                       bool                 `json:"new"`
	Status                    string               `json:"status"`
	CancelOnly                bool                 `json:"cancel_only"`
	LimitOnly                 bool                 `json:"limit_only"`
	PostOnly                  bool                 `json:"post_only"`
	TradingDisabled           bool                 `json:"trading_disabled"`
	AuctionMode               bool                 `json:"auction_mode"`
	ProductType               string               `json:"product_type"`
	QuoteCurrencyId           string               `json:"quote_currency_id"`
	BaseCurrencyId            string               `json:"base_currency_id"`
	FcmSessionDetails         SessionDetails       `json:"fcm_trading_session_details"`
	MidMarketPrice            string               `json:"mid_market_price"`
	Alias                     string               `json:"alias"`
	AliasTo                   []string             `json:"alias_to"`
	BaseDisplaySymbol         string               `json:"base_display_symbol"`
	QuoteDisplaySymbol        string               `json:"quote_display_symbol"`
	ViewOnly                  bool                 `json:"view_only"`
	PriceIncrement            string               `json:"price_increment"`
	FutureProductDetails      FutureProductDetails `json:"future_product_details"`
}

type Candle struct {
	Start  string `json:"start"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
}

type Trade struct {
	TradeId   string    `json:"trade_id"`
	ProductId string    `json:"product_id"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Time      time.Time `json:"time"`
	Side      string    `json:"side"`
	Bid       string    `json:"bid"`
	Ask       string    `json:"ask"`
}

type SessionDetails struct {
	IsSessionOpen string `json:"is_session_open"`
	OpenTime      string `json:"open_time"`
	CloseTime     string `json:"close_time"`
}

type FutureProductDetails struct {
	Venue                  string           `json:"venue"`
	ContractCode           string           `json:"contract_code"`
	ContractExpiry         string           `json:"contract_expiry"`
	ContractSize           string           `json:"contract_size"`
	ContractRootUnit       string           `json:"contract_root_unit"`
	GroupDescription       string           `json:"group_description"`
	ContractExpiryTimezone string           `json:"contract_expiry_timezone"`
	GroupShortDescription  string           `json:"group_short_description"`
	RiskManagedBy          string           `json:"risk_managed_by"`
	ContractExpiryType     string           `json:"contract_expiry_type"`
	PerpetualDetails       PerpetualDetails `json:"perpetual_details"`
	ContractDisplayName    string           `json:"contract_display_name"`
}

type FeeSummary struct {
	TotalVolume             int     `json:"total_volume"`
	TotalFees               float64 `json:"total_fees"`
	FeeTier                 FeeTier `json:"fee_tier"`
	MarginRate              Rate    `json:"margin_rate"`
	GoodsAndServicesTax     Gst     `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume int     `json:"advanced_trade_only_volume"`
	AdvancedTradeOnlyFees   float64 `json:"advanced_trade_only_fees"`
	CoinbaseProVolume       int     `json:"coinbase_pro_volume"`
	CoinbaseProFees         float64 `json:"coinbase_pro_fees"`
}

type FeeTier struct {
	PricingTier  string `json:"pricing_tier"`
	UsdFrom      string `json:"usd_from"`
	UsdTo        string `json:"usd_to"`
	TakerFeeRate string `json:"taker_fee_rate"`
	MakerFeeRate string `json:"maker_fee_rate"`
	AopFrom      string `json:"aop_from"`
	AopTo        string `json:"aop_to"`
}

type Gst struct {
	Rate string `json:"rate"`
	Type string `json:"type"`
}

type PerpetualDetails struct {
	OpenInterest string `json:"open_interest"`
	FundingRate  string `json:"funding_rate"`
	FundingTime  string `json:"funding_time"`
}

type PaginationParams struct {
	Cursor string `json:"cursor"`
	Limit  string `json:"limit"`
}

type Order struct {
	OrderId               string             `json:"order_id"`
	ProductId             string             `json:"product_id"`
	UserId                string             `json:"user_id"`
	OrderConfiguration    OrderConfiguration `json:"order_configuration"`
	Side                  string             `json:"side"`
	ClientOrderId         string             `json:"client_order_id"`
	Status                string             `json:"status"`
	TimeInForce           string             `json:"time_in_force"`
	CreatedTime           string             `json:"created_time"`
	CompletionPercentage  string             `json:"completion_percentage"`
	FilledSize            string             `json:"filled_size"`
	AverageFilledPrice    string             `json:"average_filled_price"`
	NumberOfFills         string             `json:"number_of_fills"`
	FilledValue           string             `json:"filled_value"`
	PendingCancel         bool               `json:"pending_cancel"`
	SizeInQuote           bool               `json:"size_in_quote"`
	TotalFees             string             `json:"total_fees"`
	SizeInclusiveOfFees   bool               `json:"size_inclusive_of_fees"`
	TotalValueAfterFees   string             `json:"total_value_after_fees"`
	TriggerStatus         string             `json:"trigger_status"`
	OrderType             string             `json:"order_type"`
	RejectReason          string             `json:"reject_reason"`
	Settled               bool               `json:"settled"`
	ProductType           string             `json:"product_type"`
	RejectMessage         string             `json:"reject_message"`
	CancelMessage         string             `json:"cancel_message"`
	OrderPlacementSource  string             `json:"order_placement_source"`
	OutstandingHoldAmount string             `json:"outstanding_hold_amount"`
	IsLiquidation         bool               `json:"is_liquidation"`
	LastFillTime          string             `json:"last_fill_time"`
	EditHistory           []EditHistoryItem  `json:"edit_history"`
}

type OrderConfiguration struct {
	MarketMarketIoc       *MarketIoc    `json:"market_market_ioc,omitempty"`
	SorLimitIoc           *SorLimitIoc  `json:"sor_limit_ioc,omitempty"`
	LimitLimitGtc         *LimitGtc     `json:"limit_limit_gtc,omitempty"`
	LimitLimitGtd         *LimitGtd     `json:"limit_limit_gtd,omitempty"`
	LimitLimitFok         *LimitFok     `json:"limit_limit_fok,omitempty"`
	StopLimitStopLimitGtc *StopLimitGtc `json:"stop_limit_stop_limit_gtc,omitempty"`
	StopLimitStopLimitGtd *StopLimitGtd `json:"stop_limit_stop_limit_gtd,omitempty"`
	TriggerBracketGtc     *TriggerGtc   `json:"trigger_bracket_gtc,omitempty"`
	TriggerBracketGtd     *TriggerGtd   `json:"trigger_bracket_gtd,omitempty"`
}

type MarketIoc struct {
	QuoteSize string `json:"quote_size,omitempty"`
	BaseSize  string `json:"base_size,omitempty"`
}

type SorLimitIoc struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
}

type LimitGtc struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
	PostOnly   bool   `json:"post_only"`
}

type LimitGtd struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
	EndTime    string `json:"end_time"`
	PostOnly   bool   `json:"post_only"`
}

type LimitFok struct {
	BaseSize   string `json:"base_size"`
	LimitPrice string `json:"limit_price"`
}

type StopLimitGtc struct {
	BaseSize      string `json:"base_size"`
	LimitPrice    string `json:"limit_price"`
	StopPrice     string `json:"stop_price"`
	StopDirection string `json:"stop_direction"`
}

type StopLimitGtd struct {
	BaseSize      string `json:"base_size"`
	LimitPrice    string `json:"limit_price"`
	StopPrice     string `json:"stop_price"`
	EndTime       string `json:"end_time"`
	StopDirection string `json:"stop_direction"`
}

type TriggerGtc struct {
	BaseSize         string `json:"base_size"`
	LimitPrice       string `json:"limit_price"`
	StopTriggerPrice string `json:"stop_trigger_price"`
}

type TriggerGtd struct {
	BaseSize         string `json:"base_size"`
	LimitPrice       string `json:"limit_price"`
	StopTriggerPrice string `json:"stop_trigger_price"`
	EndTime          string `json:"end_time"`
}

type EditHistoryItem struct {
	Price                  string `json:"price"`
	Size                   string `json:"size"`
	ReplaceAcceptTimestamp string `json:"replace_accept_timestamp"`
}

type Pagination struct {
	HasNext bool   `json:"has_next"`
	Cursor  string `json:"cursor"`
	Size    string `json:"size"`
}

type Preview struct {
	OrderTotal       string   `json:"order_total"`
	CommissionTotal  string   `json:"commission_total"`
	Errs             []string `json:"errs"`
	Warning          []string `json:"warning"`
	QuoteSize        string   `json:"quote_size"`
	BaseSize         string   `json:"base_size"`
	BestBid          string   `json:"best_bid"`
	BestAsk          string   `json:"best_ask"`
	IsMax            string   `json:"is_max"`
	OrderMarginTotal string   `json:"order_margin_total"`
	Leverage         string   `json:"leverage"`
	LongLeverage     string   `json:"long_leverage"`
	ShortLeverage    string   `json:"short_leverage"`
	Slippage         string   `json:"slippage"`
}

type SuccessResponse struct {
	OrderId       string `json:"order_id"`
	ProductId     string `json:"product_id"`
	Side          string `json:"side"`
	ClientOrderId string `json:"client_order_id"`
}

type ErrorResponse struct {
	Error                 string `json:"error"`
	Message               string `json:"message"`
	ErrorDetails          string `json:"error_details"`
	PreviewFailureReason  string `json:"preview_failure_reason"`
	NewOrderFailureReason string `json:"new_order_failure_reason"`
}

type Fill struct {
	EntryId            string    `json:"entry_id"`
	TradeId            string    `json:"trade_id"`
	OrderId            string    `json:"order_id"`
	TradeTime          time.Time `json:"trade_time"`
	TradeType          string    `json:"trade_type"`
	Price              string    `json:"price"`
	Size               string    `json:"size"`
	Commission         string    `json:"commission"`
	ProductId          string    `json:"product_id"`
	SequenceTimestamp  time.Time `json:"sequence_timestamp"`
	LiquidityIndicator string    `json:"liquidity_indicator"`
	SizeInQuote        bool      `json:"size_in_quote"`
	UserId             string    `json:"user_id"`
	Side               string    `json:"side"`
}

type BreakdownResponse struct {
	Breakdown Breakdown `json:"breakdown"`
}

type Breakdown struct {
	Portfolio         Portfolio         `json:"portfolio"`
	PortfolioBalances PortfolioBalances `json:"portfolio_balances"`
	SpotPositions     []SpotPosition    `json:"spot_positions"`
	PerpPositions     []PerpPosition    `json:"perp_positions"`
	FuturesPositions  []FuturesPosition `json:"futures_positions"`
}

type PortfolioBalances struct {
	TotalBalance               Amount `json:"total_balance"`
	TotalFuturesBalance        Amount `json:"total_futures_balance"`
	TotalCashEquivalentBalance Amount `json:"total_cash_equivalent_balance"`
	TotalCryptoBalance         Amount `json:"total_crypto_balance"`
	FuturesUnrealizedPnl       Amount `json:"futures_unrealized_pnl"`
	PerpUnrealizedPnl          Amount `json:"perp_unrealized_pnl"`
}

type SpotPosition struct {
	Asset                string  `json:"asset"`
	AccountUuid          string  `json:"account_uuid"`
	TotalBalanceFiat     float64 `json:"total_balance_fiat"`
	TotalBalanceCrypto   float64 `json:"total_balance_crypto"`
	AvailableToTradeFiat float64 `json:"available_to_trade_fiat"`
	Allocation           float64 `json:"allocation"`
	OneDayChange         float64 `json:"one_day_change"`
	CostBasis            Amount  `json:"cost_basis"`
	AssetImgUrl          string  `json:"asset_img_url"`
	IsCash               bool    `json:"is_cash"`
}

type PerpPosition struct {
	ProductId             string            `json:"product_id"`
	ProductUuid           string            `json:"product_uuid"`
	Symbol                string            `json:"symbol"`
	AssetImageUrl         string            `json:"asset_image_url"`
	Vwap                  DualCurrencyValue `json:"vwap"`
	PositionSide          string            `json:"position_side"`
	NetSize               string            `json:"net_size"`
	BuyOrderSize          string            `json:"buy_order_size"`
	SellOrderSize         string            `json:"sell_order_size"`
	ImContribution        string            `json:"im_contribution"`
	UnrealizedPnl         DualCurrencyValue `json:"unrealized_pnl"`
	MarkPrice             DualCurrencyValue `json:"mark_price"`
	LiquidationPrice      DualCurrencyValue `json:"liquidation_price"`
	Leverage              string            `json:"leverage"`
	ImNotional            DualCurrencyValue `json:"im_notional"`
	MmNotional            DualCurrencyValue `json:"mm_notional"`
	PositionNotional      DualCurrencyValue `json:"position_notional"`
	MarginType            string            `json:"margin_type"`
	LiquidationBuffer     string            `json:"liquidation_buffer"`
	LiquidationPercentage string            `json:"liquidation_percentage"`
}

type DualCurrencyValue struct {
	UserNativeCurrency Amount `json:"userNativeCurrency"`
	RawCurrency        Amount `json:"rawCurrency"`
}

type FuturesPosition struct {
	ProductId       string `json:"product_id"`
	ContractSize    string `json:"contract_size"`
	Side            string `json:"side"`
	Amount          string `json:"amount"`
	AvgEntryPrice   string `json:"avg_entry_price"`
	CurrentPrice    string `json:"current_price"`
	UnrealizedPnl   string `json:"unrealized_pnl"`
	Expiry          string `json:"expiry"`
	UnderlyingAsset string `json:"underlying_asset"`
	AssetImgUrl     string `json:"asset_img_url"`
	ProductName     string `json:"product_name"`
	Venue           string `json:"venue"`
	NotionalValue   string `json:"notional_value"`
}

type BalanceSummary struct {
	FuturesBuyingPower          Amount `json:"futures_buying_power"`
	TotalUsdBalance             Amount `json:"total_usd_balance"`
	CbiUsdBalance               Amount `json:"cbi_usd_balance"`
	CfmUsdBalance               Amount `json:"cfm_usd_balance"`
	TotalOpenOrdersHoldAmount   Amount `json:"total_open_orders_hold_amount"`
	UnrealizedPnl               Amount `json:"unrealized_pnl"`
	DailyRealizedPnl            Amount `json:"daily_realized_pnl"`
	InitialMargin               Amount `json:"initial_margin"`
	AvailableMargin             Amount `json:"available_margin"`
	LiquidationThreshold        Amount `json:"liquidation_threshold"`
	LiquidationBufferAmount     Amount `json:"liquidation_buffer_amount"`
	LiquidationBufferPercentage string `json:"liquidation_buffer_percentage"`
}

type Rate struct {
	Value string `json:"value"`
}

type CancelResult struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failure_reason"`
	OrderId       string `json:"order_id"`
}

type EditError struct {
	EditFailureReason    string `json:"edit_failure_reason"`
	PreviewFailureReason string `json:"preview_failure_reason"`
}

type TradeIncentiveMetadata struct {
	UserIncentiveId string `json:"user_incentive_id"`
	CodeVal         string `json:"code_val"`
}

type Convert struct {
	Id                 string           `json:"id"`
	Status             string           `json:"status"`
	UserEnteredAmount  Amount           `json:"user_entered_amount"`
	Amount             Amount           `json:"amount"`
	Subtotal           Amount           `json:"subtotal"`
	Total              Amount           `json:"total"`
	Fees               []Fee            `json:"fees"`
	TotalFee           Fee              `json:"total_fee"`
	Source             SourceTargetInfo `json:"source"`
	Target             SourceTargetInfo `json:"target"`
	UnitPrice          UnitPriceInfo    `json:"unit_price"`
	UserWarnings       []UserWarning    `json:"user_warnings"`
	UserReference      string           `json:"user_reference"`
	SourceCurrency     string           `json:"source_currency"`
	TargetCurrency     string           `json:"target_currency"`
	CancellationReason ErrorInfo        `json:"cancellation_reason"`
	SourceId           string           `json:"source_id"`
	TargetId           string           `json:"target_id"`
	ExchangeRate       Amount           `json:"exchange_rate"`
	TaxDetails         []TaxDetail      `json:"tax_details"`
	TradeIncentiveInfo TradeIncentive   `json:"trade_incentive_info"`
	TotalFeeWithoutTax Fee              `json:"total_fee_without_tax"`
	FiatDenotedTotal   Amount           `json:"fiat_denoted_total"`
}

type Fee struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Amount      Amount     `json:"amount"`
	Label       string     `json:"label"`
	Disclosure  Disclosure `json:"disclosure"`
}

type Disclosure struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        Link   `json:"link"`
}

type Link struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type SourceTargetInfo struct {
	Type          string        `json:"type"`
	Network       string        `json:"network"`
	LedgerAccount LedgerAccount `json:"ledger_account"`
}

type LedgerAccount struct {
	AccountId string `json:"account_id"`
	Currency  string `json:"currency"`
	Owner     Owner  `json:"owner"`
}

type Owner struct {
	Id       string `json:"id"`
	Uuid     string `json:"uuid"`
	UserUuid string `json:"user_uuid"`
	Type     string `json:"type"`
}

type UnitPriceInfo struct {
	TargetToFiat   Amount `json:"target_to_fiat"`
	TargetToSource Amount `json:"target_to_source"`
	SourceToFiat   Amount `json:"source_to_fiat"`
}

type UserWarning struct {
	Id      string  `json:"id"`
	Link    Link    `json:"link"`
	Context Context `json:"context"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
}

type Context struct {
	Details  []string `json:"details"`
	Title    string   `json:"title"`
	LinkText string   `json:"link_text"`
}

type ErrorInfo struct {
	Message   string `json:"message"`
	Code      string `json:"code"`
	ErrorCode string `json:"error_code"`
	ErrorCta  string `json:"error_cta"`
}

type TaxDetail struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type TradeIncentive struct {
	AppliedIncentive    bool   `json:"applied_incentive"`
	UserIncentiveId     string `json:"user_incentive_id"`
	CodeVal             string `json:"code_val"`
	EndsAt              string `json:"ends_at"`
	FeeWithoutIncentive Amount `json:"fee_without_incentive"`
	Redeemed            bool   `json:"redeemed"`
}

type PaymentMethod struct {
	Id            string    `json:"id"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Currency      string    `json:"currency"`
	Verified      bool      `json:"verified"`
	AllowBuy      bool      `json:"allow_buy"`
	AllowSell     bool      `json:"allow_sell"`
	AllowDeposit  bool      `json:"allow_deposit"`
	AllowWithdraw bool      `json:"allow_withdraw"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Sweep struct {
	Id              string `json:"id"`
	RequestedAmount Amount `json:"requested_amount"`
	ShouldSweepAll  bool   `json:"should_sweep_all"`
	Status          string `json:"status"`
	ScheduledTime   string `json:"scheduled_time"`
}

type Schema struct {
	Type string `json:"type"`
}

type IntxPortfolio struct {
	PortfolioUuid              string       `json:"portfolio_uuid"`
	Collateral                 string       `json:"collateral"`
	PositionNotional           string       `json:"position_notional"`
	OpenPositionNotional       string       `json:"open_position_notional"`
	PendingFees                string       `json:"pending_fees"`
	Borrow                     string       `json:"borrow"`
	AccruedInterest            string       `json:"accrued_interest"`
	RollingDebt                string       `json:"rolling_debt"`
	PortfolioInitialMargin     string       `json:"portfolio_initial_margin"`
	PortfolioImNotional        Amount       `json:"portfolio_im_notional"`
	PortfolioMaintenanceMargin string       `json:"portfolio_maintenance_margin"`
	PortfolioMmNotional        Amount       `json:"portfolio_mm_notional"`
	LiquidationPercentage      string       `json:"liquidation_percentage"`
	LiquidationBuffer          string       `json:"liquidation_buffer"`
	MarginType                 string       `json:"margin_type"`
	MarginFlags                string       `json:"margin_flags"`
	LiquidationStatus          string       `json:"liquidation_status"`
	UnrealizedPnl              Amount       `json:"unrealized_pnl"`
	TotalBalance               Amount       `json:"total_balance"`
	Summary                    PerpsSummary `json:"summary"`
}

type PerpsSummary struct {
	UnrealizedPnl       Amount `json:"unrealized_pnl"`
	BuyingPower         Amount `json:"buying_power"`
	TotalBalance        Amount `json:"total_balance"`
	MaxWithdrawalAmount Amount `json:"max_withdrawal_amount"`
}

type IntxPosition struct {
	ProductId        string `json:"product_id"`
	ProductUuid      string `json:"product_uuid"`
	PortfolioUuid    string `json:"portfolio_uuid"`
	Symbol           string `json:"symbol"`
	Vwap             Amount `json:"vwap"`
	EntryVwap        Amount `json:"entry_vwap"`
	PositionSide     string `json:"position_side"`
	MarginType       string `json:"margin_type"`
	NetSize          string `json:"net_size"`
	BuyOrderSize     string `json:"buy_order_size"`
	SellOrderSize    string `json:"sell_order_size"`
	ImContribution   string `json:"im_contribution"`
	UnrealizedPnl    Amount `json:"unrealized_pnl"`
	MarkPrice        Amount `json:"mark_price"`
	LiquidationPrice Amount `json:"liquidation_price"`
	Leverage         string `json:"leverage"`
	ImNotional       Amount `json:"im_notional"`
	MmNotional       Amount `json:"mm_notional"`
	PositionNotional string `json:"position_notional"`
}
type IntxSummary struct {
	AggregatedPnl Amount `json:"aggregated_pnl"`
}
