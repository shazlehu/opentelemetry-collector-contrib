// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledProvidedByUser bool
}

// IsEnabledProvidedByUser returns true if `enabled` option is explicitly set in user settings to any value.
func (ms *MetricSettings) IsEnabledProvidedByUser() bool {
	return ms.enabledProvidedByUser
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for apachereceiver metrics.
type MetricsSettings struct {
	ApacheCPULoad            MetricSettings `mapstructure:"apache.cpu.load"`
	ApacheCPUTime            MetricSettings `mapstructure:"apache.cpu.time"`
	ApacheCurrentConnections MetricSettings `mapstructure:"apache.current_connections"`
	ApacheLoad1              MetricSettings `mapstructure:"apache.load.1"`
	ApacheLoad15             MetricSettings `mapstructure:"apache.load.15"`
	ApacheLoad5              MetricSettings `mapstructure:"apache.load.5"`
	ApacheRequestTime        MetricSettings `mapstructure:"apache.request.time"`
	ApacheRequests           MetricSettings `mapstructure:"apache.requests"`
	ApacheScoreboard         MetricSettings `mapstructure:"apache.scoreboard"`
	ApacheTraffic            MetricSettings `mapstructure:"apache.traffic"`
	ApacheUptime             MetricSettings `mapstructure:"apache.uptime"`
	ApacheWorkers            MetricSettings `mapstructure:"apache.workers"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ApacheCPULoad: MetricSettings{
			Enabled: true,
		},
		ApacheCPUTime: MetricSettings{
			Enabled: true,
		},
		ApacheCurrentConnections: MetricSettings{
			Enabled: true,
		},
		ApacheLoad1: MetricSettings{
			Enabled: true,
		},
		ApacheLoad15: MetricSettings{
			Enabled: true,
		},
		ApacheLoad5: MetricSettings{
			Enabled: true,
		},
		ApacheRequestTime: MetricSettings{
			Enabled: true,
		},
		ApacheRequests: MetricSettings{
			Enabled: true,
		},
		ApacheScoreboard: MetricSettings{
			Enabled: true,
		},
		ApacheTraffic: MetricSettings{
			Enabled: true,
		},
		ApacheUptime: MetricSettings{
			Enabled: true,
		},
		ApacheWorkers: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeCPULevel specifies the a value cpu_level attribute.
type AttributeCPULevel int

const (
	_ AttributeCPULevel = iota
	AttributeCPULevelSelf
	AttributeCPULevelChildren
)

// String returns the string representation of the AttributeCPULevel.
func (av AttributeCPULevel) String() string {
	switch av {
	case AttributeCPULevelSelf:
		return "self"
	case AttributeCPULevelChildren:
		return "children"
	}
	return ""
}

// MapAttributeCPULevel is a helper map of string to AttributeCPULevel attribute value.
var MapAttributeCPULevel = map[string]AttributeCPULevel{
	"self":     AttributeCPULevelSelf,
	"children": AttributeCPULevelChildren,
}

// AttributeCPUMode specifies the a value cpu_mode attribute.
type AttributeCPUMode int

const (
	_ AttributeCPUMode = iota
	AttributeCPUModeSystem
	AttributeCPUModeUser
)

// String returns the string representation of the AttributeCPUMode.
func (av AttributeCPUMode) String() string {
	switch av {
	case AttributeCPUModeSystem:
		return "system"
	case AttributeCPUModeUser:
		return "user"
	}
	return ""
}

// MapAttributeCPUMode is a helper map of string to AttributeCPUMode attribute value.
var MapAttributeCPUMode = map[string]AttributeCPUMode{
	"system": AttributeCPUModeSystem,
	"user":   AttributeCPUModeUser,
}

// AttributeScoreboardState specifies the a value scoreboard_state attribute.
type AttributeScoreboardState int

const (
	_ AttributeScoreboardState = iota
	AttributeScoreboardStateOpen
	AttributeScoreboardStateWaiting
	AttributeScoreboardStateStarting
	AttributeScoreboardStateReading
	AttributeScoreboardStateSending
	AttributeScoreboardStateKeepalive
	AttributeScoreboardStateDnslookup
	AttributeScoreboardStateClosing
	AttributeScoreboardStateLogging
	AttributeScoreboardStateFinishing
	AttributeScoreboardStateIdleCleanup
	AttributeScoreboardStateUnknown
)

// String returns the string representation of the AttributeScoreboardState.
func (av AttributeScoreboardState) String() string {
	switch av {
	case AttributeScoreboardStateOpen:
		return "open"
	case AttributeScoreboardStateWaiting:
		return "waiting"
	case AttributeScoreboardStateStarting:
		return "starting"
	case AttributeScoreboardStateReading:
		return "reading"
	case AttributeScoreboardStateSending:
		return "sending"
	case AttributeScoreboardStateKeepalive:
		return "keepalive"
	case AttributeScoreboardStateDnslookup:
		return "dnslookup"
	case AttributeScoreboardStateClosing:
		return "closing"
	case AttributeScoreboardStateLogging:
		return "logging"
	case AttributeScoreboardStateFinishing:
		return "finishing"
	case AttributeScoreboardStateIdleCleanup:
		return "idle_cleanup"
	case AttributeScoreboardStateUnknown:
		return "unknown"
	}
	return ""
}

// MapAttributeScoreboardState is a helper map of string to AttributeScoreboardState attribute value.
var MapAttributeScoreboardState = map[string]AttributeScoreboardState{
	"open":         AttributeScoreboardStateOpen,
	"waiting":      AttributeScoreboardStateWaiting,
	"starting":     AttributeScoreboardStateStarting,
	"reading":      AttributeScoreboardStateReading,
	"sending":      AttributeScoreboardStateSending,
	"keepalive":    AttributeScoreboardStateKeepalive,
	"dnslookup":    AttributeScoreboardStateDnslookup,
	"closing":      AttributeScoreboardStateClosing,
	"logging":      AttributeScoreboardStateLogging,
	"finishing":    AttributeScoreboardStateFinishing,
	"idle_cleanup": AttributeScoreboardStateIdleCleanup,
	"unknown":      AttributeScoreboardStateUnknown,
}

// AttributeWorkersState specifies the a value workers_state attribute.
type AttributeWorkersState int

const (
	_ AttributeWorkersState = iota
	AttributeWorkersStateBusy
	AttributeWorkersStateIdle
)

// String returns the string representation of the AttributeWorkersState.
func (av AttributeWorkersState) String() string {
	switch av {
	case AttributeWorkersStateBusy:
		return "busy"
	case AttributeWorkersStateIdle:
		return "idle"
	}
	return ""
}

// MapAttributeWorkersState is a helper map of string to AttributeWorkersState attribute value.
var MapAttributeWorkersState = map[string]AttributeWorkersState{
	"busy": AttributeWorkersStateBusy,
	"idle": AttributeWorkersStateIdle,
}

type metricApacheCPULoad struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.cpu.load metric with initial data.
func (m *metricApacheCPULoad) init() {
	m.data.SetName("apache.cpu.load")
	m.data.SetDescription("Current load of the CPU.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
}

func (m *metricApacheCPULoad) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheCPULoad) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheCPULoad) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheCPULoad(settings MetricSettings) metricApacheCPULoad {
	m := metricApacheCPULoad{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheCPUTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.cpu.time metric with initial data.
func (m *metricApacheCPUTime) init() {
	m.data.SetName("apache.cpu.time")
	m.data.SetDescription("Jiffs used by processes of given category.")
	m.data.SetUnit("{jiff}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, cpuLevelAttributeValue string, cpuModeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("level", cpuLevelAttributeValue)
	dp.Attributes().PutStr("mode", cpuModeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheCPUTime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheCPUTime(settings MetricSettings) metricApacheCPUTime {
	m := metricApacheCPUTime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheCurrentConnections struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.current_connections metric with initial data.
func (m *metricApacheCurrentConnections) init() {
	m.data.SetName("apache.current_connections")
	m.data.SetDescription("The number of active connections currently attached to the HTTP server.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricApacheCurrentConnections) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheCurrentConnections) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheCurrentConnections) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheCurrentConnections(settings MetricSettings) metricApacheCurrentConnections {
	m := metricApacheCurrentConnections{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheLoad1 struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.load.1 metric with initial data.
func (m *metricApacheLoad1) init() {
	m.data.SetName("apache.load.1")
	m.data.SetDescription("The average server load during the last minute.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
}

func (m *metricApacheLoad1) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheLoad1) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheLoad1) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheLoad1(settings MetricSettings) metricApacheLoad1 {
	m := metricApacheLoad1{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheLoad15 struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.load.15 metric with initial data.
func (m *metricApacheLoad15) init() {
	m.data.SetName("apache.load.15")
	m.data.SetDescription("The average server load during the last 15 minutes.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
}

func (m *metricApacheLoad15) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheLoad15) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheLoad15) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheLoad15(settings MetricSettings) metricApacheLoad15 {
	m := metricApacheLoad15{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheLoad5 struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.load.5 metric with initial data.
func (m *metricApacheLoad5) init() {
	m.data.SetName("apache.load.5")
	m.data.SetDescription("The average server load during the last 5 minutes.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
}

func (m *metricApacheLoad5) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheLoad5) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheLoad5) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheLoad5(settings MetricSettings) metricApacheLoad5 {
	m := metricApacheLoad5{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheRequestTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.request.time metric with initial data.
func (m *metricApacheRequestTime) init() {
	m.data.SetName("apache.request.time")
	m.data.SetDescription("Total time spent on handling requests.")
	m.data.SetUnit("ms")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricApacheRequestTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheRequestTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheRequestTime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheRequestTime(settings MetricSettings) metricApacheRequestTime {
	m := metricApacheRequestTime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheRequests struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.requests metric with initial data.
func (m *metricApacheRequests) init() {
	m.data.SetName("apache.requests")
	m.data.SetDescription("The number of requests serviced by the HTTP server per second.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricApacheRequests) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheRequests) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheRequests(settings MetricSettings) metricApacheRequests {
	m := metricApacheRequests{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheScoreboard struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.scoreboard metric with initial data.
func (m *metricApacheScoreboard) init() {
	m.data.SetName("apache.scoreboard")
	m.data.SetDescription("The number of workers in each state.")
	m.data.SetUnit("{workers}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheScoreboard) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, scoreboardStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", scoreboardStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheScoreboard) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheScoreboard) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheScoreboard(settings MetricSettings) metricApacheScoreboard {
	m := metricApacheScoreboard{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheTraffic struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.traffic metric with initial data.
func (m *metricApacheTraffic) init() {
	m.data.SetName("apache.traffic")
	m.data.SetDescription("Total HTTP server traffic.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricApacheTraffic) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheTraffic) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheTraffic) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheTraffic(settings MetricSettings) metricApacheTraffic {
	m := metricApacheTraffic{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheUptime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.uptime metric with initial data.
func (m *metricApacheUptime) init() {
	m.data.SetName("apache.uptime")
	m.data.SetDescription("The amount of time that the server has been running in seconds.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricApacheUptime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheUptime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheUptime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheUptime(settings MetricSettings) metricApacheUptime {
	m := metricApacheUptime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheWorkers struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.workers metric with initial data.
func (m *metricApacheWorkers) init() {
	m.data.SetName("apache.workers")
	m.data.SetDescription("The number of workers currently attached to the HTTP server.")
	m.data.SetUnit("{workers}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheWorkers) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, workersStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", workersStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheWorkers) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheWorkers) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheWorkers(settings MetricSettings) metricApacheWorkers {
	m := metricApacheWorkers{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                int                 // maximum observed number of metrics per resource.
	resourceCapacity               int                 // maximum observed number of resource attributes.
	metricsBuffer                  pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo // contains version information
	metricApacheCPULoad            metricApacheCPULoad
	metricApacheCPUTime            metricApacheCPUTime
	metricApacheCurrentConnections metricApacheCurrentConnections
	metricApacheLoad1              metricApacheLoad1
	metricApacheLoad15             metricApacheLoad15
	metricApacheLoad5              metricApacheLoad5
	metricApacheRequestTime        metricApacheRequestTime
	metricApacheRequests           metricApacheRequests
	metricApacheScoreboard         metricApacheScoreboard
	metricApacheTraffic            metricApacheTraffic
	metricApacheUptime             metricApacheUptime
	metricApacheWorkers            metricApacheWorkers
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      buildInfo,
		metricApacheCPULoad:            newMetricApacheCPULoad(settings.ApacheCPULoad),
		metricApacheCPUTime:            newMetricApacheCPUTime(settings.ApacheCPUTime),
		metricApacheCurrentConnections: newMetricApacheCurrentConnections(settings.ApacheCurrentConnections),
		metricApacheLoad1:              newMetricApacheLoad1(settings.ApacheLoad1),
		metricApacheLoad15:             newMetricApacheLoad15(settings.ApacheLoad15),
		metricApacheLoad5:              newMetricApacheLoad5(settings.ApacheLoad5),
		metricApacheRequestTime:        newMetricApacheRequestTime(settings.ApacheRequestTime),
		metricApacheRequests:           newMetricApacheRequests(settings.ApacheRequests),
		metricApacheScoreboard:         newMetricApacheScoreboard(settings.ApacheScoreboard),
		metricApacheTraffic:            newMetricApacheTraffic(settings.ApacheTraffic),
		metricApacheUptime:             newMetricApacheUptime(settings.ApacheUptime),
		metricApacheWorkers:            newMetricApacheWorkers(settings.ApacheWorkers),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithApacheServerName sets provided value as "apache.server.name" attribute for current resource.
func WithApacheServerName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("apache.server.name", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/apachereceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricApacheCPULoad.emit(ils.Metrics())
	mb.metricApacheCPUTime.emit(ils.Metrics())
	mb.metricApacheCurrentConnections.emit(ils.Metrics())
	mb.metricApacheLoad1.emit(ils.Metrics())
	mb.metricApacheLoad15.emit(ils.Metrics())
	mb.metricApacheLoad5.emit(ils.Metrics())
	mb.metricApacheRequestTime.emit(ils.Metrics())
	mb.metricApacheRequests.emit(ils.Metrics())
	mb.metricApacheScoreboard.emit(ils.Metrics())
	mb.metricApacheTraffic.emit(ils.Metrics())
	mb.metricApacheUptime.emit(ils.Metrics())
	mb.metricApacheWorkers.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordApacheCPULoadDataPoint adds a data point to apache.cpu.load metric.
func (mb *MetricsBuilder) RecordApacheCPULoadDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheCPULoad, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCPULoad.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheCPUTimeDataPoint adds a data point to apache.cpu.time metric.
func (mb *MetricsBuilder) RecordApacheCPUTimeDataPoint(ts pcommon.Timestamp, inputVal string, cpuLevelAttributeValue AttributeCPULevel, cpuModeAttributeValue AttributeCPUMode) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheCPUTime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCPUTime.recordDataPoint(mb.startTime, ts, val, cpuLevelAttributeValue.String(), cpuModeAttributeValue.String())
	return nil
}

// RecordApacheCurrentConnectionsDataPoint adds a data point to apache.current_connections metric.
func (mb *MetricsBuilder) RecordApacheCurrentConnectionsDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheCurrentConnections, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCurrentConnections.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheLoad1DataPoint adds a data point to apache.load.1 metric.
func (mb *MetricsBuilder) RecordApacheLoad1DataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad1, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad1.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheLoad15DataPoint adds a data point to apache.load.15 metric.
func (mb *MetricsBuilder) RecordApacheLoad15DataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad15, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad15.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheLoad5DataPoint adds a data point to apache.load.5 metric.
func (mb *MetricsBuilder) RecordApacheLoad5DataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for ApacheLoad5, value was %s: %w", inputVal, err)
	}
	mb.metricApacheLoad5.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheRequestTimeDataPoint adds a data point to apache.request.time metric.
func (mb *MetricsBuilder) RecordApacheRequestTimeDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheRequestTime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheRequestTime.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheRequestsDataPoint adds a data point to apache.requests metric.
func (mb *MetricsBuilder) RecordApacheRequestsDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheRequests, value was %s: %w", inputVal, err)
	}
	mb.metricApacheRequests.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheScoreboardDataPoint adds a data point to apache.scoreboard metric.
func (mb *MetricsBuilder) RecordApacheScoreboardDataPoint(ts pcommon.Timestamp, val int64, scoreboardStateAttributeValue AttributeScoreboardState) {
	mb.metricApacheScoreboard.recordDataPoint(mb.startTime, ts, val, scoreboardStateAttributeValue.String())
}

// RecordApacheTrafficDataPoint adds a data point to apache.traffic metric.
func (mb *MetricsBuilder) RecordApacheTrafficDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricApacheTraffic.recordDataPoint(mb.startTime, ts, val)
}

// RecordApacheUptimeDataPoint adds a data point to apache.uptime metric.
func (mb *MetricsBuilder) RecordApacheUptimeDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheUptime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheUptime.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordApacheWorkersDataPoint adds a data point to apache.workers metric.
func (mb *MetricsBuilder) RecordApacheWorkersDataPoint(ts pcommon.Timestamp, inputVal string, workersStateAttributeValue AttributeWorkersState) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheWorkers, value was %s: %w", inputVal, err)
	}
	mb.metricApacheWorkers.recordDataPoint(mb.startTime, ts, val, workersStateAttributeValue.String())
	return nil
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
